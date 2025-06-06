package main

import (
	"fmt"
	"sync"
)

const limit = 10

var (
	oddsPrinter = func(wg *sync.WaitGroup, oddsCh chan struct{}, evensCh chan struct{}) {
		defer wg.Done()
		limit := limit - 1
		for i := 1; i <= limit; i += 2 {
			<-oddsCh
			fmt.Println("print odd", i)

			if i == limit {
				close(evensCh)

				return
			}

			evensCh <- struct{}{}
		}
	}

	evensPrinter = func(wg *sync.WaitGroup, oddsCh chan struct{}, evensCh chan struct{}) {
		defer wg.Done()
		for i := 2; i <= limit; i += 2 {
			<-evensCh
			fmt.Println("print even", i)

			if i == limit {
				close(oddsCh)

				return
			}

			oddsCh <- struct{}{}
		}
	}
)

func main() {
	oddsCh := make(chan struct{})
	evensCh := make(chan struct{})

	wg := &sync.WaitGroup{}
	//nolint:mnd
	wg.Add(2)

	go oddsPrinter(wg, oddsCh, evensCh)
	go evensPrinter(wg, oddsCh, evensCh)

	oddsCh <- struct{}{}

	wg.Wait()

	fmt.Println("done")
}
