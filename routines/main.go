package main

import (
	"echoStars/routines/concurrencyType"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls []string

func main() {
	fmt.Println("Routines")
	urls = append(urls, "www.google.com", "https://www.github.com", "www.bitbucket.org", "https://www.hexacta.com", "https://go.dev/", "https://www.npm.com")
	getSites(concurrencyType.Sequential)
	getSites(concurrencyType.Sequential)
	getSites(concurrencyType.ConcurrentWG)
	getSites(concurrencyType.ConcurrentChan)
}

func getSites(concurrencyType concurrencyType.ConcurrencyType) {
	start := time.Now()
	var err []error
	switch concurrencyType {
	case 0:
		err = sequentialGet()
		fmt.Println("Sequential get took: ", time.Since(start))
	case 1:
		err = concurrentWaitGroupGet()
		fmt.Println("Concurrent waitGroup get took: ", time.Since(start))
	case 2:
		err = concurrentGetChn()
		fmt.Println("Concurrent channel get took: ", time.Since(start))
	}

	if err != nil {
		fmt.Println("Error list: ", err)
	}
}

func sequentialGet() []error {
	fmt.Println("Sequential get init")
	var errs []error
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			errs = append(errs, fmt.Errorf("error in seguential get: %w", err))
		} else {
			fmt.Println("//GET STATUS RES: ", res.StatusCode, res.Request.URL)
		}
	}

	return errs
}

func concurrentWaitGroupGet() []error {
	fmt.Println("Concurrent get with WaitGroup init")
	var waitGroup sync.WaitGroup
	var errs []error
	for _, url := range urls {
		waitGroup.Add(1)
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				errs = append(errs, fmt.Errorf("error in concurrent waitgroup get: %w", err))
			} else {
				fmt.Println("//GET STATUS RES: ", res.StatusCode, res.Request.URL)
			}
			waitGroup.Done()
		}(url)
	}
	waitGroup.Wait()
	return errs
}

func concurrentGetChn() []error {
	fmt.Println("Concurrent get with Channels init")
	errChan := make(chan error)
	done := make(chan int)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				err = fmt.Errorf("error in concurrent channel get: %w", err)
				errChan <- err
			} else {
				fmt.Println("//GET STATUS RES: ", res.StatusCode, res.Request.URL)
				done <- res.StatusCode
			}
		}(url)
	}
	count := 0
	var errList []error
	for {
		count++
		select {
		case err := <-errChan:
			errList = append(errList, err)
		case statusCode := <-done:
			fmt.Println("Status code: ", statusCode)
			if count == len(urls) {
				return errList
			}
		}
	}
}
