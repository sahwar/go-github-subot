package main

import (
	"encoding/json"
	"testing"
)

func setCredentials() {
	config.UserName = "Subot"
	// Write your token here before running test
	config.Token = "7d5fe99ab4e0b00f887ea166ecef16e8333f07e8"
}

func TestUsers(t *testing.T) {
	setCredentials()
	var users []User

	queries := make(map[string]string)
	queries["per_page"] = "1"

	data, status := get(API+"/users", queries)
	// Attempt to parse response as array of User structure
	if err := json.Unmarshal(data, &users); err != nil || status != SUCCESS {
		t.Fatal(err, status, string(data))
	}

	// The first user. Should be mojombo
	mojombo := users[0]

	// Gives an error message if user isn't mojombo
	if mojombo.Id != 1 || mojombo.Login != "mojombo" {
		t.Fatal(mojombo)
	}
}

func TestFollowing(t *testing.T) {
	setCredentials()
	var users []User

	queries := make(map[string]string)
	queries["per_page"] = "1"

	data, status := get(API+"/user/following", queries)
	// Attempt to parse response as array of User structure
	if err := json.Unmarshal(data, &users); err != nil || status != SUCCESS {
		t.Fatal(err, status, string(data))
	}

	// The first user. Should be mojombo
	mojombo := users[0]

	// Gives an error message if user isn't mojombo
	if mojombo.Id != 1 || mojombo.Login != "mojombo" {
		t.Fatal(mojombo)
	}
}

func TestFollowers(t *testing.T) {
	setCredentials()
	var users []User

	queries := make(map[string]string)
	queries["per_page"] = "1"

	data, status := get(API+"/user/followers", queries)
	// Attempt to parse response as array of User structure
	if err := json.Unmarshal(data, &users); err != nil || status != SUCCESS {
		t.Fatal(err, status, string(data))
	}

	// The first user. Should be hea955
	hea := users[0]

	// Gives an error message if user isn't hea955
	if hea.Id != 48537272 || hea.Login != "hea955" {
		t.Fatal(hea)
	}
}
