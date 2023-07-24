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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(c, url)
	}

	// for i := 0; i < len(urls); i++ {
	for u := range c {
		go func(url string) {
			time.Sleep(time.Second * 2)
			checkUrl(c, url)
		}(u)

	}
}

func checkUrl(c chan string, url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down! :o")
		c <- url
		return
	}

	fmt.Println(url, "is up! :)")
	c <- url
}
