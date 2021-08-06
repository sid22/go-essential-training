package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
 "user": "harry p",
 "type": "deposit",
 "amount": 10.0
}
`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data)

	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error can't decode %s\n", err)
	}

	fmt.Printf("got the val %+v\n", req)

	prevBal := 12.0
	resp := map[string]interface{}{ // interface{} allows to keep any data type
		"ok":      true,
		"balance": prevBal + req.Amount,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error in enc %s\n", err)
	}

}
