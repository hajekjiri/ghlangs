package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"math"
	"os"
	"sort"
	"strings"
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

func listReposWithLanguages(repos *[]repoEntry, unit string) {
	for i, repo := range *repos {
		fmt.Printf("%s\n", repo.nameWithOwner)

		listLanguages(&repo.langs, "bytes", "descending", unit)

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

func listLanguages(langs *[]langEntry, sortKey string, sortDirection string, unit string) {
	var sortFunc func(a, b int) bool
	switch sortKey {
	case "name":
		switch sortDirection {
		case "ascending":
			sortFunc = func(a, b int) bool {
				return strings.Compare((*langs)[a].name, (*langs)[b].name) < 0
			}
		case "descending":
			sortFunc = func(a, b int) bool {
				return strings.Compare((*langs)[a].name, (*langs)[b].name) > 0
			}
		default:
			log.Printf(
				"Warning: unknown sort direction '%s' in listLanguages(), defaulting to 'ascending'\n",
				sortDirection,
			)
			sortFunc = func(a, b int) bool {
				return strings.Compare((*langs)[a].name, (*langs)[b].name) < 0
			}
		}
	case "bytes":
		switch sortDirection {
		case "ascending":
			sortFunc = func(a, b int) bool {
				return (*langs)[a].bytes < (*langs)[b].bytes
			}
		case "descending":
			sortFunc = func(a, b int) bool {
				return (*langs)[a].bytes > (*langs)[b].bytes
			}
		default:
			log.Printf(
				"Warning: unknown sort direction '%s' in listLanguages(), defaulting to descending\n",
				sortDirection,
			)
			sortFunc = func(a, b int) bool {
				return (*langs)[a].bytes > (*langs)[b].bytes
			}
		}
	case "":
		sortFunc = nil
	default:
		log.Printf("Warning: unknown sort key '%s' in listLanguages()\n", sortKey)
		sortFunc = nil
	}

	if sortFunc != nil {
		sort.Slice(*langs, sortFunc)
	}

	totalSize := 0
	for _, lang := range *langs {
		totalSize += lang.bytes
	}

	totalSizeLabel := "Total size"
	totalSizeString := getSizeByUnit(totalSize, unit)
	maxNameLen := len(totalSizeLabel)
	maxSizeLen := len(totalSizeString)
	for _, lang := range *langs {
		if len(lang.name) > maxNameLen {
			maxNameLen = len(lang.name)
		}

		langSizeLen := len(getSizeByUnit(lang.bytes, unit))
		if langSizeLen > maxSizeLen {
			maxSizeLen = langSizeLen
		}
	}

	dashes := make([]byte, len(totalSizeLabel)+maxSizeLen+11)
	for i := range dashes {
		dashes[i] = '-'
	}
	dashesStr := string(dashes)

	fmt.Println(dashesStr)
	fmt.Printf("|%s|%s|100.00%%|\n", totalSizeLabel, strlpad(totalSizeString, maxSizeLen))
	fmt.Println(dashesStr)

	for _, lang := range *langs {
		relativeSize := float64(lang.bytes) / float64(totalSize) * 100
		fmt.Printf(
			"|%s|%s|%s%%|\n",
			strrpad(lang.name, maxNameLen),
			strlpad(getSizeByUnit(lang.bytes, unit), maxSizeLen),
			strlpad(fmt.Sprintf("%.2f", relativeSize), len("100.00")),
		)
	}
	fmt.Println(dashesStr)
}

func strlpad(str string, pad int) string {
	if pad < len(str) {
		return string(str)
	}

	whitespace := make([]byte, pad-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

	bytes := []byte(str)
	bytes = append(whitespace, bytes...)
	return string(bytes)
}

func strrpad(str string, pad int) string {
	if pad < len(str) {
		return string(str)
	}

	whitespace := make([]byte, pad-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

	bytes := []byte(str)
	bytes = append(bytes, whitespace...)
	return string(bytes)
}

func getSizeByUnit(size int, unit string) string {
	var exp int
	switch unit {
	case "auto":
		return getAutoSize(size)
	case "B":
		exp = 1
	case "kB":
		exp = -3
	case "MB":
		exp = -6
	case "GB":
		exp = -9
	case "TB":
		exp = -12
	case "PB":
		exp = -15
	case "EB":
		exp = -18
	// no need for more units because 10^18 approaches the limits of 64bit integers
	default:
		log.Printf("Warning: unknown unit '%s' in getSizeByUnit(), defaulting to B\n", unit)
		unit = " B"
		exp = 1
	}

	return fmt.Sprintf("%.3f %s", float64(size)*math.Pow10(exp), unit)
}

func getAutoSize(size int) string {
	unitNo := 0
	var unit string
	sizeFloat := float64(size)
	for sizeFloat > 1000 {
		sizeFloat = sizeFloat / 1000
		unitNo++
	}

	switch unitNo {
	case 0:
		unit = " B"
	case 1:
		unit = "kB"
	case 2:
		unit = "MB"
	case 3:
		unit = "GB"
	case 4:
		unit = "TB"
	case 5:
		unit = "PB"
	case 6:
		unit = "EB"
		// no need for more units because 10^18 approaches the limits of 64bit integers
	default:
		log.Fatal("Error in getAutoSize(): this shouldn't have happened because 64bit integers can't reach sizes larger than ~10^18")
	}

	return fmt.Sprintf("%.3f %s", sizeFloat, unit)
}

type queryFirstRepos struct {
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
		} `graphql:"repositories(first: 100, affiliations: OWNER)"`
	}
	RateLimit struct {
		Limit     int
		Cost      int
		Remaining int
		ResetAt   string
	}
}

type queryNextRepos struct {
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
		} `graphql:"repositories(first: 100, affiliations: OWNER, after: $after)"`
	}
	RateLimit struct {
		Limit     int
		Cost      int
		Remaining int
		ResetAt   string
	}
}

func getRepos(client *githubv4.Client) (*[]repoEntry, error) {
	query := queryFirstRepos{}
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

	fmt.Printf(
		"Progress: %d/%d repositories (API Rate Limit %d/%d)\n",
		len(query.Viewer.Repositories.Nodes),
		query.Viewer.Repositories.TotalCount,
		query.RateLimit.Limit-query.RateLimit.Remaining,
		query.RateLimit.Limit,
	)

	if query.Viewer.Repositories.PageInfo.HasNextPage {
		getNextRepos(client, repos, query.Viewer.Repositories.PageInfo.EndCursor,
			len(query.Viewer.Repositories.Nodes))
	}

	return repos, nil
}

func getNextRepos(client *githubv4.Client, repos *[]repoEntry, offset string, progress int) error {
	query := queryNextRepos{}
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

	fmt.Printf(
		"Progress: %d/%d repositories (API Rate Limit %d/%d)\n",
		progress+len(query.Viewer.Repositories.Nodes),
		query.Viewer.Repositories.TotalCount,
		query.RateLimit.Limit-query.RateLimit.Remaining,
		query.RateLimit.Limit,
	)

	if query.Viewer.Repositories.PageInfo.HasNextPage {
		getNextRepos(client, repos, query.Viewer.Repositories.PageInfo.EndCursor, progress+len(query.Viewer.Repositories.Nodes))
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

	sortKey := "bytes"
	sortDirection := "descending"
	unit := "auto"

	langs := getLanguagesFromRepos(repos)
	listReposWithLanguages(repos, unit)
	fmt.Println()
	fmt.Println("All repositories:")
	listLanguages(langs, sortKey, sortDirection, unit)
}
