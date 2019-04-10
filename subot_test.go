package main

import (
	"encoding/json"
	"testing"
)

func TestSubot(t *testing.T) {
	// Sets username to send as User-Agent
	config.UserName = "Subot"
	// Sets token to work as authorized user
	config.Token = "ee01c5067e591928deca6c902da04be48f89f9fa"
	var users []User

	// Get response from https://api.github.com/users
	data, status := get(API+"/users", nil)
	// Attempt to parse response as array of User structure
	if err := json.Unmarshal(data, &users); err != nil || status != 200 {
		t.Fatal(err, status, string(data))
	}

	// The first user. Should be mojombo
	mojombo := users[0]

	// Gives an error message if user isn't mojombo
	if mojombo.Id != 1 || mojombo.Login != "mojombo" {
		t.Fatal(mojombo)
	}
}
