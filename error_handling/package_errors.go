// print a comment here
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open configuration file")
	}

	defer file.Close()
	cfg := &Config{}
	// parse file and put it in struct
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("configpath/aconf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s \n", err)
		log.Panicf("error %+v \n", err)
		os.Exit(1)
	}

	fmt.Println(cfg)

	cfg2, err := readConfig("configpath/aconf2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s \n", err)
		log.Panicf("error %+v \n", err)
		os.Exit(1)
	}

	fmt.Println(cfg2)
}
