// print a comment here
package main

import (
	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidfile string) error {
	data, err := ioutil.ReadFile(pidfile)
	if err != nil {
		return errors.Wrap(err, "can't open pid file")
	}

	err2 := os.Remove(pidfile)
	if err2 != nil {
		log.Printf("warning, can;t remove pidfile")
	}

	stpid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(stpid)
	if err != nil {
		return errors.Wrap(err, "unable to convert prcess to int")
	}

	fmt.Printf("kill process %d \n ", pid)
	return nil
}

func main() {
	err := killServer("pidfile")
	if err != nil {
		fmt.Println(err)
	}
}
