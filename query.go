package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
)

func getRepos(client *githubv4.Client) ([]repoEntry, error) {
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
			} `graphql:"repositories(first: 100, affiliations: OWNER, after: $after)"`
		}
		RateLimit struct {
			Limit     int
			Cost      int
			Remaining int
			ResetAt   string
		}
	}

	repos := []repoEntry{}
	var offset *string = nil
	progress := 0
	query.Viewer.Repositories.PageInfo.HasNextPage = true
	for query.Viewer.Repositories.PageInfo.HasNextPage {
		params := map[string]interface{}{
			"after": (*githubv4.String)(offset),
		}
		err := client.Query(context.Background(), &query, params)
		if err != nil {
			return []repoEntry{}, err
		}

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

		progress += len(query.Viewer.Repositories.Nodes)
		offset = &query.Viewer.Repositories.PageInfo.EndCursor

		fmt.Printf(
			"Progress: %d/%d repositories (API Rate Limit %d/%d)\n",
			progress,
			query.Viewer.Repositories.TotalCount,
			query.RateLimit.Limit-query.RateLimit.Remaining,
			query.RateLimit.Limit,
		)
	}

	return repos, nil
}
