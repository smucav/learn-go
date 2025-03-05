// Project to practice for
// go routines
// channels
// mutex
// non blocking channel operation
// custom errors
// struct
// interface
// generic types
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// interface for coin fetcher
type PriceFetcher interface {
	FetchPrice(coin string) (float64, error)
	Name() string
}

// fetch coin from coin gecko site
type CoinGeckoFetcher struct{}

func (cg *CoinGeckoFetcher) FetchPrice(coin string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", coin)
	resp, err := http.Get(url)
	if err != nil {
		return 0, &FetchError{Source: cg.Name(), Reason: err.Error()}
	}

	// it's return JSON look like {"ETH":
	//                                   {"usd": 2189.03}
	//                            }
	var data map[string]map[string]float64

	defer resp.Body.Close()

	// check for the error while assigning JSON response to data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, &FetchError{Source: cg.Name(), Reason: "Failed to parse JSON: " + err.Error()}
	}

	price, ok := data[coin]["usd"]

	if !ok {
		return 0, &FetchError{Source: cg.Name(), Reason: "Price not found for " + coin}
	}
	return price, nil
}

// Gets it's name to know which site we are using to fetch crypto value
// because we are using PriceFetcher interface
func (cg *CoinGeckoFetcher) Name() string {
	return "CoinGecko"
}

// cryptocompare is open and free but require api key
type CryptoCompareFetcher struct {
	apiKey string
}

func (cd *CryptoCompareFetcher) FetchPrice(coin string) (float64, error) {
	// map bitcoin to BTC because crypto compare uses 'BTC' instead of bitcoin which is
	// 'bitcoin' on coin gecko
	coin_map := map[string]string{
		"bitcoin":  "BTC",
		"ethereum": "ETH",
	}
	coin = coin_map[coin]

	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD&api_key=%s", coin, cd.apiKey)
	resp, err := http.Get(url)

	if err != nil {
		return 0, &FetchError{Source: cd.Name(), Reason: err.Error()}
	}

	// put interface to achieve generic type because JSON returns string but function return type is
	// float64
	var data map[string]interface{}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, &FetchError{Source: cd.Name(), Reason: "Failed to parse JSON" + err.Error()}
	}

	price, ok := data["USD"]
	if !ok {
		return 0, &FetchError{Source: cd.Name(), Reason: ""}
	}
	if p, ok := price.(float64); ok {
		return p, nil
	}
	return 0, &FetchError{Source: cd.Name(), Reason: "Price is not a number"}
}

func (cd *CryptoCompareFetcher) Name() string {
	return "CoinDesk"
}

// use custom error message
type FetchError struct {
	Source string
	Reason string
}

func (e *FetchError) Error() string {
	return fmt.Sprintf("Fetch from %s Failed %s\n", e.Source, e.Reason)
}

// cache the current fetched price
type PriceCache[T any] struct {
	mu    sync.Mutex
	cache map[string]T
}

// create new price cache using a function to make it simple
func NewPriceCache[T any]() *PriceCache[T] {
	return &PriceCache[T]{cache: make(map[string]T)}
}

// use key value pair to store cache value ex: CoinDesk-bitcoin: 88643.00
func (pc *PriceCache[T]) Store(key string, value T) {
	defer pc.mu.Unlock()
	pc.mu.Lock()
	pc.cache[key] = value
}

func (pc *PriceCache[T]) Load(key string) (T, bool) {
	defer pc.mu.Unlock()
	pc.mu.Lock()
	value, ok := pc.cache[key]
	return value, ok
}

// define a type to send result through channel
type FetchResponse struct {
	Source string
	Coin   string
	Price  float64
	Err    error
}

func main() {
	// create buffered channel to avoid any deadlock and block
	fetchCh := make(chan FetchResponse, 2)
	statusCh := make(chan string)

	cache := NewPriceCache[float64]()
	apiKey := os.Getenv("CRYPTOCOMPARE_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: CRYPTOCOMPARE_API_KEY environment variable not set")
		return
	}

	sites := []PriceFetcher{
		&CoinGeckoFetcher{},
		&CryptoCompareFetcher{apiKey: apiKey},
	}

	coins := []string{"bitcoin", "ethereum"}

	var wg sync.WaitGroup
	for _, site := range sites {
		for _, coin := range coins {
			wg.Add(1)
			go func(ch chan<- FetchResponse, coins string) {
				defer wg.Done()
				price, err := site.FetchPrice(coin)
				ch <- FetchResponse{
					Source: site.Name(),
					Coin:   coin,
					Price:  price,
					Err:    err,
				}
			}(fetchCh, coin)
			select {
			case statusCh <- fmt.Sprintf("%s Fetched %s", site.Name(), coin):
			default:
			}
		}
	}
	// fetch the result according to crypto values we fetch
	// to each site we fetch btc and eth
	TotFetches := len(sites) * len(coins)

	for i := 0; i < TotFetches; i++ {
		select {
		case resp := <-fetchCh:
			if resp.Err != nil {
				fmt.Printf("Error from: %s for %s: %v\n", resp.Source, resp.Coin, resp.Err)
			} else {
				// store to cache once we get from the channel
				cache.Store(fmt.Sprintf("%s-%s", resp.Source, resp.Coin), resp.Price)
				fmt.Printf("Cached %s from %s: $%.2f\n", resp.Coin, resp.Source, resp.Price)
			}
		case <-time.After(2 * time.Second):
			fmt.Printf("Timeout: some fetches are too slow\n")
			i--
		case status := <-statusCh:
			fmt.Println("Status: ", status)
		default:
			fmt.Println("No activity at this moment, checking again...")
			time.Sleep(2 * time.Second)
			i--
		}
	}

	// wait for all go routine finished
	wg.Wait()
	// close the channel after that
	close(fetchCh)

	// thread safe cache
	cache.mu.Lock()
	fmt.Println("\nFinal Price Cache: ")
	for key, price := range cache.cache {
		fmt.Printf("%s: $%.2f\n", key, price)
	}
	cache.mu.Unlock()
}
