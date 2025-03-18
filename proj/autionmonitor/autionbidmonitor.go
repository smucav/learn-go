package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type StockFetcher interface {
	FetchPrice(symbol string) (float64, error)
	Name() string
}

type FetchError struct {
	Source string
	Reason string
}

func (fe *FetchError) Error() string {
	return fmt.Sprintf("From %s Failed: %s\n", fe.Source, fe.Reason)
}

type MockStock struct {
	prices map[string]float64
	mu     sync.Mutex
}

func (ms *MockStock) FetchPrice(symbol string) (float64, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	price, err := ms.prices[symbol]
	if !err {
		return 0, &FetchError{Source: ms.Name(), Reason: "Can't find symbol"}
	}
	price += float64(time.Now().Nanosecond()%10) / 10
	return price, nil
}

func (ms *MockStock) Name() string {
	return "Mock Stock"
}

type PriceCache[T any] struct {
	mu    sync.Mutex
	cache map[string]T
}

func NewPriceCache[T any]() *PriceCache[T] {
	return &PriceCache[T]{cache: make(map[string]T)}
}

func (pc *PriceCache[T]) Store(key string, value T) {
	pc.mu.Lock()
	pc.mu.Unlock()
	pc.cache[key] = value
}

type Bid struct {
	Symbol string
	Amount float64
}

type Auction struct {
	fetcher    StockFetcher
	currentBid Bid
	bidCh      chan Bid
	cache      *PriceCache[float64]
	mu         sync.Mutex
}

func NewAuction(fetcher StockFetcher) *Auction {
	return &Auction{
		fetcher: fetcher,
		bidCh:   make(chan Bid, 10),
		cache:   NewPriceCache[float64](),
	}
}

func (a *Auction) Run(wg *sync.WaitGroup, signal chan<- bool) {
	fmt.Println("Auction started..")
	sym := a.currentBid.Symbol
	fmt.Printf("Current bid $%.2f of %s\n", a.currentBid.Amount, sym)
	a.cache.Store(sym, a.currentBid.Amount)
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(15 * time.Second)
	for {

		select {
		case bid := <-a.bidCh:
			a.mu.Lock()
			if bid.Amount > a.currentBid.Amount {
				fmt.Printf("New bid $%.2f\n", bid.Amount)
				a.cache.Store(sym, bid.Amount)
				a.currentBid.Amount = bid.Amount
			}
			a.mu.Unlock()
			if !timer.Stop() {
				<-timer.C
			} else {
				timer.Reset(15 * time.Second)
			}
		case <-ticker.C:
			switch v := a.fetcher.(type) {
			case *MockStock:
				v.prices[sym] = a.currentBid.Amount
			default:
				fmt.Println("no match")
			}
			price, err := a.fetcher.FetchPrice(sym)
			if err != nil {
				fmt.Printf("Error fetching %s\n", sym)
				continue
			}

			fmt.Printf("New price for %s is $%.2f\n", sym, price)
		case <-timer.C:
			signal <- true
			close(signal)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	fetcher := &MockStock{
		prices: map[string]float64{"AAPL": 230.01, "GOOGL": 150.02},
	}
	auction := NewAuction(fetcher)
	sym := "AAPL"
	price, err := fetcher.FetchPrice(sym)
	signal := make(chan bool, 1)

	if err != nil {
		fmt.Printf("Error Fetching %s - %s\n", fetcher.Name(), err.Error())
		return
	}

	auction.currentBid = Bid{Symbol: sym, Amount: price}

	wg.Add(1)
	go auction.Run(&wg, signal)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Connection error: ", err)
		return
	}
	wg.Add(1)
	go HandleConnection(conn, &wg, auction, signal)

	wg.Wait()
}

func HandleConnection(conn net.Conn, wg *sync.WaitGroup, auction *Auction, signal <-chan bool) {
	defer wg.Done()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	inputCh := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				close(inputCh)
				return
			}
			inputCh <- strings.TrimSpace(input)
		}
	}()

	for {
		select {
		case <-signal:
			fmt.Printf("Bidding Ended.... winner $%.2f for %s\n", auction.currentBid.Amount, auction.currentBid.Symbol)
			fmt.Fprintln(conn, "EXIT")
			return
		case input, ok := <-inputCh:
			if !ok {
				fmt.Println("Input channel closed")
				os.Exit(1)
				return
			}
			if input == "exit" {
				return
			}
			bid, err := strconv.Atoi(input)
			if err != nil {
				fmt.Fprintf(conn, "Invalid bid: %v\n", err)
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				auction.bidCh <- Bid{Symbol: "AAPL", Amount: float64(bid)}
			}()
		}
	}
}
