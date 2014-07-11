package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
)

var client *http.Client

type StuffToPost struct {
	Name string
	Username string
	email string
}

func makeRequest(i int) *http.Request {
	stuff := StuffToPost{"Stewart Duffey", "sjduffey", "sduffey@inviqa.com"}
	encodedStuff, _ := json.Marshal(stuff)

	r, _ := http.NewRequest(
		"POST", 
		"http://localhost:8080/?count=" + strconv.Itoa(i), 
		bytes.NewBuffer(encodedStuff))
	r.Header.Add("Content-Type", "application/json")

	return r
}

func pushRequest(c chan<- *http.Request) {
	for i := 0; i < 10; i++ {
		c <- makeRequest(i)
	}
}

func popRequest(c <-chan *http.Request) {
	for {
		req := <- c

		client := &http.Client{}
		resp, _ := client.Do(req)

		printResponseBody(resp)
	}
}

func printResponseBody(r *http.Response) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body[:]))
}

func main() {

	client = &http.Client{}

	var c chan *http.Request = make(chan *http.Request, 3)

	go pushRequest(c)
	go popRequest(c)

	var scan string
	fmt.Scanln(&scan)
}
