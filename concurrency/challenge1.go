package main

import (
	"fmt"
	"net/http"
	// not the recommended way to sync b/w go routines
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

	chanl := make(chan int)
	for i, url := range urls {
		go func(url string, i int) {
			returnType(url)
			chanl <- i
		}(url, i)
	}

	for range urls {
		fmt.Printf("got %d\n", <-chanl)
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	// not the recommended way to sync b/w go routines
// )

// func returnType(url string, out chan string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		out <- fmt.Sprintf("error is &s\n", err)
// 		return
// 	}

// 	defer resp.Body.Close()
// 	ctype := resp.Header.Get("Content-Type")
// 	out <- fmt.Sprintf("%s >> %s\n", url, ctype)
// }

// func main() {
// 	urls := []string{
// 		"https://www.rubrik.com",
// 		"https://api.github.com",
// 		"https://httpbin.org/xml",
// 	}

// 	chanl := make(chan string)
// 	for _, url := range urls {
// 		go returnType(url, chanl)
// 	}

// 	for range urls {
// 		fmt.Printf("got %s \n", <-chanl)
// 	}
// }
