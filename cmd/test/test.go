package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	Concurrency = 25
	Requests    = 1000
)

func main() {
	var (
		counter uint32
		success uint32
		failed  uint32

		wg sync.WaitGroup
	)

	client := &http.Client{}

	worker := func(i int, requests uint32, wg *sync.WaitGroup) {
		fmt.Printf("Starting the worker #%d\n", i)
		defer fmt.Println("Stopping the worker #%d\n", i)
		defer wg.Done()

		var j uint32

		for j = 0; j <= requests; j++ {

			wait := rand.Intn(100) // let's wait a little bit before the next test
			time.Sleep(time.Duration(wait) * time.Millisecond)

			req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/ping", nil)
			resp, err := client.Do(req)

			if err != nil {
				atomic.AddUint32(&failed, 1)

				continue
			}

			if resp.StatusCode == http.StatusOK {
				atomic.AddUint32(&success, 1)
			} else {
				atomic.AddUint32(&failed, 1)
			}

			atomic.AddUint32(&counter, 1)

			resp.Body.Close()
		}
	}

	for i := 0; i < Concurrency; i++ {
		wg.Add(1)
		go worker(i, Requests/Concurrency, &wg)
	}

	wg.Wait()

	fmt.Printf("Requests done: %d\n", Requests)
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Failed: %d\n", failed)
}
