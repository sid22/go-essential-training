package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.rubrik.com/")
	if err != nil {
		log.Fatalf("err in get %s\n", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
