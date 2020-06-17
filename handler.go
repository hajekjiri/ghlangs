package main

func getLanguagesFromRepos(repos []repoEntry) []langEntry {
	langMap := make(map[string]int)
	langSlice := make([]langEntry, 0)
	for _, repo := range repos {
		for _, lang := range repo.langs {
			langMap[lang.name] += lang.size
		}
	}

	for lang, size := range langMap {
		langSlice = append(langSlice, langEntry{lang, size})
	}

	return langSlice
}

type reposQuery interface {
	GetRepositories() repoResult
	GetRateLimit() rateLimitResult
}

func extractReposFromQuery(q reposQuery) []repoEntry {
	repos := []repoEntry{}
	queryRepos := q.GetRepositories()
	for _, repo := range queryRepos.Nodes {
		e := repoEntry{
			nameWithOwner: repo.NameWithOwner,
			langs:         []langEntry{},
		}
		for _, lang := range repo.Languages.Edges {
			e.langs = append(e.langs, langEntry{lang.Node.Name, lang.Size})
		}
		repos = append(repos, e)
	}

	return repos
}
