package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
)

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

func getRepos(client *githubv4.Client) ([]repoEntry, error) {
	query := queryFirstRepos{}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		return []repoEntry{}, err
	}

	repos := []repoEntry{}
	for _, repo := range query.Viewer.Repositories.Nodes {
		e := repoEntry{
			nameWithOwner: repo.NameWithOwner,
			langs:         []langEntry{},
		}
		for _, lang := range repo.Languages.Edges {
			e.langs = append(e.langs, langEntry{lang.Node.Name, lang.Size})
		}
		repos = append(repos, e)
	}

	fmt.Printf(
		"Progress: %d/%d repositories (API Rate Limit %d/%d)\n",
		len(query.Viewer.Repositories.Nodes),
		query.Viewer.Repositories.TotalCount,
		query.RateLimit.Limit-query.RateLimit.Remaining,
		query.RateLimit.Limit,
	)

	if query.Viewer.Repositories.PageInfo.HasNextPage {
		err := getNextRepos(client, &repos, query.Viewer.Repositories.PageInfo.EndCursor, len(query.Viewer.Repositories.Nodes))
		if err != nil {
			return []repoEntry{}, err
		}
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
