package main

//#region Header
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var (
	config Config
)

const (
	API        string = "https://api.github.com"
	SUCCESS    int    = 200
	NO_CONTENT int    = 204
)

type User struct {
	Login       string    `json:"login"`
	Id          int       `json:"id"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"public_gists"`
	Followers   int       `json:"followers"`
	Stars       int       `json:"stars"`
	Watchers    int       `json:"watchers"`
	Language    string    `json:"language"`
	Projects    []Project `json:"projects"`
}

type Project struct {
	Name            string `json:"name"`
	StargazersCount int    `json:"stargazers_count"`
	Watchers        int    `json:"watchers"`
	Language        string `json:"language"`
}

type Config struct {
	Token     string `json:"token"`
	UserName  string `json:"username"`
	Source    string `json:"source"`
	Page      int    `json:"page"`
	Id        int    `json:"id"`
	Followers int    `json:"followers"`
	Stars     int    `json:"stars"`
	Repos     int    `json:"repos"`
	Watchers  int    `json:"watchers"`
	Gists     int    `json:"gists"`
	Language  string `json:"language"`
}

//#endregion

//#region Net
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

//#endregion

func saveConfig() {
	// Convert "Config" structure to JSON format
	data, err := json.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	// Write settings to config.json file
	err = ioutil.WriteFile("Config.json", data, 0664)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config successfully saved!")
}

func checkAndFollow(users []User) {
	for _, u := range users {
		// Ignore Organizations
		if u.Type != "User" {
			continue
		}
		if config.Source == "all" {
			config.Id = u.Id
		}
		if config.Followers != 0 || config.Repos != 0 ||
			config.Gists != 0 || config.Watchers != 0 ||
			config.Stars != 0 || config.Language != "" {
			u.info()
			if config.Followers > u.Followers ||
				config.Repos > u.PublicRepos ||
				config.Gists > u.PublicGists {
				fmt.Println(u.Login, "doesn't meet requirements! [1]")
				continue
			}
		}
		if config.Stars != 0 || config.Language != "" || config.Watchers != 0 {
			u.repositories()
			u.Stars = 0
			u.Watchers = 0
			languages := make(map[string]int)
			for _, project := range u.Projects {
				u.Stars += project.StargazersCount
				u.Watchers += project.Watchers
				if config.Language != "" {
					if _, exists := languages[project.Language]; exists {
						languages[project.Language]++
					} else {
						languages[project.Language] = 1
					}
				}
			}
			usages := 0
			for key, value := range languages {
				if usages < value {
					usages = value
					u.Language = key
				}
			}
			if config.Stars > u.Stars ||
				config.Watchers > u.Watchers ||
				config.Language != u.Language {
				fmt.Println(u.Login, "doesn't meet requirements! [2]")
				continue
			}
		}
		u.follow(nil)
	}
}

func (user *User) repositories() {
	queries := make(map[string]string)
	queries["per_page"] = "100"

	// Get total amount of pages
	pages := int(math.Floor(float64(user.PublicRepos/100)) + 1)
	for i := 1; i <= pages; i++ {
		// Set page to current counter
		queries["page"] = strconv.Itoa(i)
		for _, repo := range user.getRepositoriesForPage(queries) {
			user.Projects = append(user.Projects, repo)
		}
	}
}

func (user *User) info() {
	data, status := get(user.URL, nil)
	if status != SUCCESS {
		fmt.Println("Error received. Status:", status, "\nMessage:", string(data))
		time.Sleep(time.Second * 90)
		user.info()
		return
	}
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
}

func (user *User) follow(queries map[string]string) {
	data, status := put(API+"/user/following/"+user.Login, nil)
	if status != NO_CONTENT {
		fmt.Println("Error received. Status:", status, "\nMessage:", string(data))
		time.Sleep(time.Second * 90)
		user.follow(queries)
		return
	}
	fmt.Println(user.Login, "with ID", user.Id, "now following!")
	time.Sleep(time.Second * 1)
}

func (user *User) getRepositoriesForPage(queries map[string]string) []Project {
	data, status := get(user.URL+"/repos", queries)
	if status != SUCCESS {
		fmt.Println("Error received. Status:", status, "\nMessage:", string(data))
		time.Sleep(time.Second * 90)
		return user.getRepositoriesForPage(queries)
	}
	var repositories []Project
	if err := json.Unmarshal(data, &repositories); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	return repositories
}

func getFollowers(queries map[string]string) []User {
	data, status := get(API+"/users/"+config.Source+"/followers", queries)
	if status != SUCCESS {
		fmt.Println("Error received. Status:", status, "\nMessage:", string(data))
		time.Sleep(time.Second * 90)
		return getFollowers(queries)
	}
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	return users
}

func getUsers(queries map[string]string) []User {
	data, status := get(API+"/users", queries)
	if status != SUCCESS {
		fmt.Println("Error received. Status:", status, "\nMessage:", string(data))
		time.Sleep(time.Second * 90)
		return getUsers(queries)
	}
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	return users
}

func main() {
	data, err := ioutil.ReadFile("Config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	defer saveConfig()

	go func() {
		// Catch interrupt signal to (stop executing smoothly?)
		// and save some data to config
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)

		// Block executing until signal
		<-signalChan

		// Save settings to file
		saveConfig()

		// Stop executing program with 0 code
		os.Exit(0)
	}()

	queries := make(map[string]string)
	queries["per_page"] = "100"

	var users []User

	if config.Source != "all" {
		if config.Page <= 0 {
			config.Page = 1
		}
		for {
			queries["page"] = strconv.Itoa(config.Page)
			users = getFollowers(queries)
			if len(users) == 0 {
				fmt.Println("All users from", config.Source, "successfully added to following list!")
				config.Source = "all"
				config.Page = 1
				return
			}
			checkAndFollow(users)
			config.Page++
		}
	} else {
		for {
			queries["since"] = strconv.Itoa(config.Id)
			users = getUsers(queries)
			checkAndFollow(users)
		}
	}

}
