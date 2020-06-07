package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
	"sort"
)

func authenticate(token string) (*githubv4.Client, error) {
	if token == "" {
		return &githubv4.Client{}, errors.New("Provided oauth token is an empty string")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)

	return githubv4.NewClient(httpClient), nil
}

type langEntry struct {
	name  string
	bytes int
}

type repoEntry struct {
	nameWithOwner string
	langs         []langEntry
}

func listReposWithLanguages(repos *[]repoEntry) {
	for i, repo := range *repos {
		fmt.Printf("Repository %s contains\n", repo.nameWithOwner)
		for _, lang := range repo.langs {
			fmt.Printf("%d bytes of %s\n", lang.bytes, lang.name)
		}
		if i < len(*repos)-1 {
			fmt.Println()
		}
	}
}

func getLanguagesFromRepos(repos *[]repoEntry) *[]langEntry {
	langMap := make(map[string]int)
	langSlice := make([]langEntry, 0)
	for _, repo := range *repos {
		for _, lang := range repo.langs {
			langMap[lang.name] += lang.bytes
		}
	}

	for lang, bytes := range langMap {
		langSlice = append(langSlice, langEntry{lang, bytes})
	}

	return &langSlice
}

func listLanguages(langs *[]langEntry, sortKey string, sortDirection string) {
	if sortKey != "" && sortDirection != "" {
		sort.Slice(*langs,
			func(a, b int) bool { return (*langs)[a].bytes > (*langs)[b].bytes })
	}

	for _, lang := range *langs {
		fmt.Printf("%s: %d bytes\n", lang.name, lang.bytes)
	}
}

func getRepos(client *githubv4.Client) (*[]repoEntry, error) {
	var query struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
				Nodes      []struct {
					NameWithOwner string
					Languages     struct {
						Edges []struct {
							Node struct {
								Name string
							}
							Size int
						}
					} `graphql:"languages(first : 100)"`
				}
				PageInfo struct {
					StartCursor string
					EndCursor   string
					HasNextPage bool
				}
			} `graphql:"repositories(first: 100)"`
		}
		RateLimit struct {
			Limit     int
			Cost      int
			Remaining int
			ResetAt   string
		}
	}

	err := client.Query(context.Background(), &query, nil)

	if err != nil {
		return &[]repoEntry{}, err
	}

	repos := &[]repoEntry{}
	for _, repo := range query.Viewer.Repositories.Nodes {
		e := repoEntry{
			nameWithOwner: repo.NameWithOwner,
			langs:         []langEntry{},
		}
		for _, lang := range repo.Languages.Edges {
			e.langs = append(e.langs, langEntry{lang.Node.Name, lang.Size})
		}
		*repos = append(*repos, e)
	}

	fmt.Printf("Found %d repositories\n", query.Viewer.Repositories.TotalCount)
	fmt.Printf("Rate limit: %d/%d\n",
		query.RateLimit.Limit-query.RateLimit.Remaining, query.RateLimit.Limit)

	if query.Viewer.Repositories.PageInfo.HasNextPage {
		getNextRepos(client, repos, query.Viewer.Repositories.PageInfo.EndCursor)
	}

	return repos, nil
}

func getNextRepos(client *githubv4.Client, repos *[]repoEntry, offset string) error {
	var query struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
				Nodes      []struct {
					NameWithOwner string
					Languages     struct {
						Edges []struct {
							Node struct {
								Name string
							}
							Size int
						}
					} `graphql:"languages(first : 100)"`
				}
				PageInfo struct {
					StartCursor string
					EndCursor   string
					HasNextPage bool
				}
			} `graphql:"repositories(first: 100, after: $after)"`
		}
		RateLimit struct {
			Limit     int
			Cost      int
			Remaining int
			ResetAt   string
		}
	}

	params := map[string]interface{}{
		"after": githubv4.String(offset),
	}
	err := client.Query(context.Background(), &query, params)

	if err != nil {
		return err
	}

	if repos == nil {
		repos = &[]repoEntry{}
	}

	for _, repo := range query.Viewer.Repositories.Nodes {
		e := repoEntry{
			nameWithOwner: repo.NameWithOwner,
			langs:         []langEntry{},
		}
		for _, lang := range repo.Languages.Edges {
			e.langs = append(e.langs, langEntry{lang.Node.Name, lang.Size})
		}
		*repos = append(*repos, e)
	}

	fmt.Printf("Rate limit: %d/%d\n",
		query.RateLimit.Limit-query.RateLimit.Remaining, query.RateLimit.Limit)

	if query.Viewer.Repositories.PageInfo.HasNextPage {
		getNextRepos(client, repos, query.Viewer.Repositories.PageInfo.EndCursor)
	}

	return nil
}

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	client, err := authenticate(token)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	repos, err := getRepos(client)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	langs := getLanguagesFromRepos(repos)
	listLanguages(langs, "bytes", "desc")
}
