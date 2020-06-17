package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"strconv"
)

func getViewerRepos(client *githubv4.Client) ([]repoEntry, error) {
	query := viewerReposQuery{}
	repos := []repoEntry{}
	var offset *string = nil
	for {
		params := map[string]interface{}{
			"after": (*githubv4.String)(offset),
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
			"Progress: %s/%d repositories (API Rate Limit %d/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			rateLimit.Limit-rateLimit.Remaining,
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}

func getUserRepos(client *githubv4.Client, login string) ([]repoEntry, error) {
	query := userReposQuery{}
	repos := []repoEntry{}
	var offset *string = nil
	for {
		params := map[string]interface{}{
			"after": (*githubv4.String)(offset),
			"login": githubv4.String(login),
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
			"Progress: %s/%d repositories (API Rate Limit %d/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			rateLimit.Limit-rateLimit.Remaining,
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}

func getOrgRepos(client *githubv4.Client, login string) ([]repoEntry, error) {
	query := orgReposQuery{}
	repos := []repoEntry{}
	var offset *string = nil
	for {
		params := map[string]interface{}{
			"after": (*githubv4.String)(offset),
			"login": githubv4.String(login),
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
			"Progress: %s/%d repositories (API Rate Limit %d/%d)\n",
			Strlpad(strconv.Itoa(len(repos)), progressPad),
			queryRepos.TotalCount,
			rateLimit.Limit-rateLimit.Remaining,
			rateLimit.Limit,
		)

		if !queryRepos.PageInfo.HasNextPage {
			break
		}
	}

	return repos, nil
}
