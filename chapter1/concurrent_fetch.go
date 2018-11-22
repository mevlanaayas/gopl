// Program that fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type FetchResult struct {
	content    string
	finishedAt int
	url        string
}

func main() {
	fmt.Println("concurrent fetch program to fetch all given urls")
	const prefix = "http://"
	resultChannel := make(chan FetchResult)
	programStart := time.Now()
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		go func(url string) {
			httpStart := time.Now()
			response, err := http.Get(url)
			if err != nil {
				fmt.Println("GET: error time: %T, error msg: %T", time.Now(), err)
			}
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("READ: error time: %T, error msg: %T", time.Now(), err)
			}
			fmt.Printf("\n url: %s, length of content: %d \n", url, len(content))
			httpElapsed := time.Since(httpStart).Seconds()
			fmt.Println("url and elapsed: ", url, httpElapsed)
			localResult := FetchResult{string(content)[:10], time.Now().Second(), url}
			resultChannel <- localResult
			defer response.Body.Close()

		}(url)
	}
	for range os.Args[1:] {
		result := <-resultChannel
		fmt.Println(result)
	}
	programElapsed := time.Since(programStart).Seconds()
	fmt.Println("elapsed time: ", programElapsed)
}
