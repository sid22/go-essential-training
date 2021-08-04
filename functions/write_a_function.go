// print a comment here

// go has a garbade collector, we don't need to manually release memory
// for other resource sockets etc we use defer
package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ctypeheader := resp.Header["Content-Type"]

	// ctypeheader2 := resp.Header.Get("Content-Type")
	// fmt.Println(ctypeheader2)
	if len(ctypeheader) == 0 {
		return "", nil
	}
	return ctypeheader[0], nil
}

func main() {
	url := "https://www.rubrik.com/"
	typ, err := contentType(url)
	if err != nil {
		fmt.Printf("error is %s \n", err)
	} else {
		fmt.Printf("type is %s \n", typ)

	}

	url2 := "https://www.wrongrubrik.com/"
	typ2, err := contentType(url2)
	if err != nil {
		fmt.Printf("error is %s \n", err)
	} else {
		fmt.Printf("type is %s \n", typ2)

	}
}
