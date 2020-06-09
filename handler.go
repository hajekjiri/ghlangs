package main

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
