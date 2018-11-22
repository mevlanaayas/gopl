// Program that fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type FetchResult struct {
	length     float64
	finishedAt int
	url        string
	order      int
}

func main() {
	fmt.Println("concurrent fetch program to fetch all given urls")

	const prefix = "http://"
	urls := os.Args[1:]
	resultChannel := make(chan FetchResult)
	// new wait group to follow whether goroutine execution is finished or not
	var wg sync.WaitGroup

	data, err := ioutil.ReadFile("urls.txt")
	if err == nil {
		fmt.Println("reading urls from file...")
		urls = strings.Split(string(data), "\n")
	}

	fmt.Println("size of urls: ", len(urls))

	programStart := time.Now()
	for index, url := range urls {
		// add concurrent event to wait group
		wg.Add(1)
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		go func(url string, index int) {
			fmt.Printf("----- %d. http get started for: %s ------\n", index, url)
			httpStart := time.Now()

			response, err := http.Get(url)
			if err != nil {
				fmt.Println("GET: error time: %T, error msg: %T", time.Now(), err)
				errorResult :=
					FetchResult{0.0, time.Now().Second(), url, index}
				resultChannel <- errorResult
				return
			}
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("READ: error time: %T, error msg: %T", time.Now(), err)
			}
			fmt.Printf("\n finished url: %s, length of content: %d", url, len(content))
			httpElapsed := time.Since(httpStart).Seconds()
			fmt.Printf("\n finished for url: %s  and elapsed time is: %f \n", url, httpElapsed)
			localResult :=
				FetchResult{float64(len(content) / 1024.0), time.Now().Second(), url, index}
			resultChannel <- localResult
			defer response.Body.Close()
			// inform Wait group about finished job
			defer wg.Done()
		}(url, index)
	}
	// wait until all the job is done
	wg.Wait()

	/*
		note that
		we used "result" channel to communicate between channels
		when we send data to channel to (that we did in goroutine above)
		and we receive this data from channel (we did in for loop below)
		It also make the program waiting till the all goroutines is ended.
		As a result:
			wg.Wait is unnecessary here. But if you remove for loop below it will br necessary :)
	*/
	for range urls {
		result := <-resultChannel
		fmt.Printf(
			"size: %f, url: %s, finished At: %d, order: %d \n",
			result.length, result.url, result.finishedAt, result.order)
	}

	programElapsed := time.Since(programStart).Seconds()
	fmt.Println("elapsed time: ", programElapsed)
}
