package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cli/go-gh"
)

func main() {
	fmt.Println("hi world, this is the gh-merge-commit-describer extension!")
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {Login string}{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	// check the login
	fmt.Printf("running as %s\n", response.Login)

	if len(os.Args) != 2 {
		fmt.Println("Usage: gh pr-merge-commits PR_NUMBER")
		os.Exit(1)
	}

	// check if the arguments are correct
	prNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Invalid PR_NUMBER")
		os.Exit(1)
	}

	fmt.Println(prNumber)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
