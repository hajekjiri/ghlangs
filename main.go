package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Provided oauth token is an empty string")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	repos, err := getRepos(client)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	sortKey := "bytes"
	sortDirection := "descending"
	unit := "auto"

	langs := getLanguagesFromRepos(repos)
	listReposWithLanguages(repos, unit)
	fmt.Println()
	fmt.Println("All repositories:")
	listLanguages(langs, sortKey, sortDirection, unit)
}
