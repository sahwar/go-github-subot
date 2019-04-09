package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	config Config
)

const (
	API string = "https://api.github.com"
)

type User struct {
	Login     string
	Id        int
	SiteAdmin bool
	Type      string
}

type Config struct {
	Token    string
	UserName string
}

func sendRequest(method string, link string, queries map[string]string) ([]byte, int) {
	// Create new client to send requests
	client := http.Client{}
	// Add timeout to prevent client hijacking
	client.Timeout = time.Second * 10
	// Create new GET request
	req, _ := http.NewRequest(method, link, nil)

	// Add queries to request
	q := req.URL.Query()
	q.Add("access_token", config.Token)
	for key, value := range queries {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	// Add header parameter to request
	req.Header.Add("User-Agent", config.UserName)

	// Execute request
	resp, err := client.Do(req)
	// Check for errors
	if err != nil {
		log.Fatal(err)
	}
	// Close connection to save resources
	defer resp.Body.Close()
	// Read all data from response to byte[]
	data, _ := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode
}

func get(link string, queries map[string]string) ([]byte, int) {
	return sendRequest("GET", link, queries)
}

func put(link string, queries map[string]string) ([]byte, int) {
	return sendRequest("PUT", link, queries)
}

func main() {
	data, err := ioutil.ReadFile("Config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	data, _ = get(API+"/users", nil)

	fmt.Println(string(data))
}
