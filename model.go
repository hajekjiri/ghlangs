package main

type langEntry struct {
	name string
	size int
}

type repoEntry struct {
	nameWithOwner string
	langs         []langEntry
}

type repoResult struct {
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
		} `graphql:"languages(first: 100)"`
	}
	PageInfo struct {
		StartCursor string
		EndCursor   string
		HasNextPage bool
	}
}

type rateLimitResult struct {
	Limit     int
	Cost      int
	Remaining int
	ResetAt   string
}

///////////////////////////////////////////////////////////////////////////////
type viewerReposQuery struct {
	Viewer struct {
		Repositories repoResult `graphql:"repositories(first: 100, ownerAffiliations: [OWNER], isFork: false, after: $after)"`
	}
	RateLimit rateLimitResult
}

func (q viewerReposQuery) GetRepositories() repoResult {
	return q.Viewer.Repositories
}

func (q viewerReposQuery) GetRateLimit() rateLimitResult {
	return q.RateLimit
}

///////////////////////////////////////////////////////////////////////////////
type userReposQuery struct {
	User struct {
		Repositories repoResult `graphql:"repositories(first: 100, ownerAffiliations: [OWNER], isFork: false, after: $after)"`
	} `graphql:"user(login: $login)"`
	RateLimit rateLimitResult
}

func (q userReposQuery) GetRepositories() repoResult {
	return q.User.Repositories
}

func (q userReposQuery) GetRateLimit() rateLimitResult {
	return q.RateLimit
}

///////////////////////////////////////////////////////////////////////////////
type orgReposQuery struct {
	Organization struct {
		Repositories repoResult `graphql:"repositories(first: 100, ownerAffiliations: [OWNER], isFork: false, after: $after)"`
	} `graphql:"organization(login: $login)"`
	RateLimit rateLimitResult
}

func (q orgReposQuery) GetRepositories() repoResult {
	return q.Organization.Repositories
}

func (q orgReposQuery) GetRateLimit() rateLimitResult {
	return q.RateLimit
}
