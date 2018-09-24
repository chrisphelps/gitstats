package main

import (
	"fmt"
	"os"
	flag "github.com/ogier/pflag"
	"github.com/chrisphelps/gitstats/github"
)

var (
	user string
)

// todo auth key and request private counts
// todo what happens if you make private request against different user
// todo oauth flow instead of api key?
// todo start walking graph
// todo get and analyze PRs from a repo
// todo how do I find all participating repos? do I need to walk the organization?


func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	user := github.GetUser(user)

	fmt.Printf("Stats for %s(%s)\n", user.Name, user.Login)

	fmt.Printf("Public Gists: %d\n", user.PublicGists)
	fmt.Printf("Public Repos: %d\n", user.PublicRepos)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}