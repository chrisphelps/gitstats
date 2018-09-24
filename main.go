package main

import (
	"fmt"
	"os"
	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

    fmt.Printf("Hello, %s\n", user)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}