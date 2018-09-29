package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	URL               string      `json:"url"`
	Name              string      `json:"name"`
	PublicRepos       int         `json:"public_repos"`
	PrivateRepos      int         `json:"owned_private_repos"`
	TotalPrivateRepos int         `json:"total_private_repos"`
	PublicGists       int         `json:"public_gists"`
	PrivateGists      int         `json:"private_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
}

func GetUser(name string, token string) User {
	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL + userEndpoint + name, nil)
	if err != nil {
		log.Fatalf("Something went wrong contstructing the request: %s\n", err)
	}

	req.Header.Add("Authorization", "token " + token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving data w/ auth: %s\n", err)
	}

	defer resp.Body.Close()

	// read the response body and handle any errors during reading.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	// create a user variable of type "User" struct to store the "Unmarshal"-ed (aka parsed JSON) data, then return the user
	var user User
	json.Unmarshal(body, &user)
	return user
}