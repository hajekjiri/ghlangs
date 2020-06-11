package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := NewClient()
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	repos, err := getRepos(client)
	if err != nil {
		log.Fatalf("Error: %s", err)
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
