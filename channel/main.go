package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://walmart.com",
		"http://asda.com",
		"http://amazon.com",
		"http://flipkart.com",
		"http://amazon.in",
		"http://errordomaintest.com",
	}

	done := make(chan bool)
	start := time.Now()
	for _, url := range urls {
		go urlStatus(url, done)
	}

	for i := 0; i < len(urls); i++ {
		if !<-done {
			fmt.Println("ERROR has occurred")
		}
	}
	// for resp := range ch {
	// 	fmt.Println(resp)
	// }
	end := time.Now()

	fmt.Println("Total Time: ", end.Sub(start))
}

func urlStatus(url string, done chan bool) {
	start := time.Now()
	_, err := http.Get(url)
	end := time.Now()
	if err != nil {
		fmt.Println(url, "is down. Response Time:", end.Sub(start))
		done <- false
		return
	}
	fmt.Println(url, "OK. Response Time:", end.Sub(start))
	done <- true
	return
}
