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
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
}

func GetUser(name string) User {
	// send GET request to GitHub API with the requested user "name"
	resp, err := http.Get(apiURL + userEndpoint + name)
	// if err occurs during GET request, then throw error and quit application
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	// Always good practice to defer closing the response body.
	// If application crashes or function finishes successfully, GO will always execute this "defer" statement
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