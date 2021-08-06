// Makign HTTP calls
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Job is a job description
type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	apiUrl := "https://api.github.com/users/" + login

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalf("err in get %s\n", err)
	}
	defer resp.Body.Close()

	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		log.Fatalf("json encode err %s\n", err)
	}
	return user, nil
}

func main() {
	login := "sid22"
	user, err := userInfo(login)
	if err != nil {
		log.Fatalf("error as %s\n", err)
	} else {
		fmt.Printf("user info is %+v\n", user)
	}

}
