// Makign HTTP calls
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	dbStore = map[string]interface{}{}
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

type CreateRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateResponse struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type FetchRequest struct {
	Key string `json:"key"`
}

type FetchResponse struct {
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}

func dbCreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	req := &CreateRequest{}

	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := &CreateResponse{}
	dbStore[req.Key] = req.Value

	resp.Result = "OK"
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Panicf("can't encode %v- %s", resp, err)
	}
}

func dbFetchHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	req := &FetchRequest{}

	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := &FetchResponse{}
	data := dbStore[req.Key]

	resp.Result = data
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Panicf("can't encode %v- %s", resp, err)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/db/create", dbCreateHandler)
	http.HandleFunc("/db/fetch", dbFetchHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
