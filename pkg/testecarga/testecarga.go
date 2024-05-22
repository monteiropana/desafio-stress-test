package testecarga

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode   int
	ResponseTime float64
}

func TesteCarga(url string, totalReq int, concurrency int) []Result {
	var results []Result
	var mutex sync.Mutex
	var wg sync.WaitGroup

	sem := make(chan struct{}, concurrency)

	for i := 0; i < totalReq; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			startTime := time.Now()
			resp, err := http.Get(url)

			var result Result
			if err != nil {
				result = Result{StatusCode: 500, ResponseTime: time.Since(startTime).Seconds()}
			} else {
				defer resp.Body.Close()
				result = Result{StatusCode: resp.StatusCode, ResponseTime: time.Since(startTime).Seconds()}
			}

			mutex.Lock()
			results = append(results, result)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	return results
}
