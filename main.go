package main

import (
	"fmt"
	"os"
	flag "github.com/ogier/pflag"
	"github.com/chrisphelps/gitstats/github"
)

var (
	user string
	token string
)

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

	gitUser := github.GetUser(user, token)

	fmt.Printf("Stats for %s(%s)\n", gitUser.Name, gitUser.Login)

	fmt.Printf("Public Gists: %d\n", gitUser.PublicGists)
	fmt.Printf("Public Repos: %d\n", gitUser.PublicRepos)

	fmt.Printf("Private Gists: %d\n", gitUser.PrivateGists)
	fmt.Printf("Private Repos: %d\n", gitUser.PrivateRepos)
	fmt.Printf("Total Private Repos: %d\n", gitUser.TotalPrivateRepos)

}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
	flag.StringVarP(&token, "token", "t", "", "Auth Token")
}