package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
	"strconv"
)

func getViewerRepos(client *graphql.Client) ([]repoEntry, error) {
	query := viewerReposQuery{}
	repos := []repoEntry{}
	var offset *string
	for {
		params := map[string]interface{}{
			"after": (*graphql.String)(offset),
		}
		err := client.Query(context.Background(), &query, params)
		if err != nil {
			return []repoEntry{}, err
		}

		newRepos := extractReposFromQuery(query)
		repos = append(repos, newRepos...)

		queryRepos := query.GetRepositories()
		offset = &queryRepos.PageInfo.EndCursor
		rateLimit := query.GetRateLimit()

		progressPad := len(strconv.Itoa(queryRepos.TotalCount))
		fmt.Printf(
			"Progress: %s/%d repositories (API Rate Limit %s/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			Strlpad(strconv.Itoa(rateLimit.Limit-rateLimit.Remaining), 4),
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}

func getUserRepos(client *graphql.Client, login string) ([]repoEntry, error) {
	query := userReposQuery{}
	repos := []repoEntry{}
	var offset *string
	for {
		params := map[string]interface{}{
			"after": (*graphql.String)(offset),
			"login": graphql.String(login),
		}
		err := client.Query(context.Background(), &query, params)
		if err != nil {
			return []repoEntry{}, err
		}

		newRepos := extractReposFromQuery(query)
		repos = append(repos, newRepos...)

		queryRepos := query.GetRepositories()
		offset = &queryRepos.PageInfo.EndCursor
		rateLimit := query.GetRateLimit()

		progressPad := len(strconv.Itoa(queryRepos.TotalCount))
		fmt.Printf(
			"Progress: %s/%d repositories (API Rate Limit %s/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			Strlpad(strconv.Itoa(rateLimit.Limit-rateLimit.Remaining), 4),
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}

func getOrgRepos(client *graphql.Client, login string) ([]repoEntry, error) {
	query := orgReposQuery{}
	repos := []repoEntry{}
	var offset *string
	for {
		params := map[string]interface{}{
			"after": (*graphql.String)(offset),
			"login": graphql.String(login),
		}
		err := client.Query(context.Background(), &query, params)
		if err != nil {
			return []repoEntry{}, err
		}

		newRepos := extractReposFromQuery(query)
		repos = append(repos, newRepos...)

		queryRepos := query.GetRepositories()
		offset = &queryRepos.PageInfo.EndCursor
		rateLimit := query.GetRateLimit()

		progressPad := len(strconv.Itoa(queryRepos.TotalCount))
		fmt.Printf(
			"Progress: %s/%d repositories (API Rate Limit %s/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			Strlpad(strconv.Itoa(rateLimit.Limit-rateLimit.Remaining), 4),
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}
