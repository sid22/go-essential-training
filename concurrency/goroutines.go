package main

import (
	"fmt"
	"net/http"
	"sync" // not the recommended way to sync b/w go routines
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error is &s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s >> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://www.rubrik.com",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		// returnType(url)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
