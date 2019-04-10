package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	config Config
)

const (
	API     string = "https://api.github.com"
	SUCCESS int    = 200
	LIMIT   int    = 429
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
	client.Timeout = time.Second * 30
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
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("User-Agent", config.UserName)

	// Execute request
	resp, err := client.Do(req)
	// Check for errors
	if err != nil {
		log.Println(err)
		return sendRequest(method, link, queries)
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

func sendPutUntilSuccess(login string) bool {
	data, status := put(API+"/user/following/"+login, nil)
	if status != 204 {
		fmt.Println(string(data), status, login)
		time.Sleep(time.Second * 90)
		return sendPutUntilSuccess(login)
	} else {
		return true
	}
}

func main() {
	data, err := ioutil.ReadFile("Config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	var users []User

	queries := make(map[string]string)
	queries["per_page"] = "100"

	i := 13882
	for {
		queries["since"] = strconv.Itoa(i)
		data, _ = get(API+"/users", queries)
		if err = json.Unmarshal(data, &users); err != nil {
			log.Fatal(err)
		}
		for _, user := range users {
			time.Sleep(time.Millisecond * 1200)
			i = user.Id
			success := false
			for !success {
				success = sendPutUntilSuccess(user.Login)
			}
			fmt.Println(user.Login, user.Id, "are followed!")
		}
	}
}
