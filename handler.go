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
