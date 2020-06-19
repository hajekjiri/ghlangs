package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
	"os"
)

// NewClient creates a new GitHub GraphQL API v4 client.
//
// Note that GitHub GraphQL API v4 requires authentication. Please save your
// API token to an environment variable called 'GITHUB_AUTH_TOKEN'.
func NewClient() (*graphql.Client, error) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return &graphql.Client{}, fmt.Errorf("API token is undefined (GITHUB_AUTH_TOKEN)")
	}

	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	return graphql.NewClient("https://api.github.com/graphql", httpClient), nil
}
