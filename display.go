package main

import (
	"fmt"
	"sort"
	"strings"
)

// ListReposWithLanguages lists repositories and used programming languages for
// each of those repositories. Repositories are sorted by name in ascending
// order and their name is displayed in the nameWithOwner format
// (owner/repository). Sizes are displayed in the provided unit. Languages
// are sorted based on the sortKey and sortOrder parameters. Uses ListLanguages
// to list programming languages.
func ListReposWithLanguages(repos []repoEntry, sortKey string, sortOrder string, unit string) error {
	repoSortFunc := func(a, b int) bool {
		return strings.Compare(repos[a].nameWithOwner, repos[b].nameWithOwner) < 0
	}
	sort.Slice(repos, repoSortFunc)

	for i, repo := range repos {
		fmt.Printf("%s\n", repo.nameWithOwner)

		err := ListLanguages(repo.langs, sortKey, sortOrder, unit)
		if err != nil {
			return err
		}

		if i < len(repos)-1 {
			fmt.Println()
		}
	}

	return nil
}

// ListLanguages lists programming languages and their sizes based on the
// provided []langEntry slice. The output is sorted based on the sortKey and
// sortOrder parameters. Sizes are displayed in the provided unit.
func ListLanguages(langs []langEntry, sortKey string, sortOrder string, unit string) error {
	var sortFunc func(a, b int) bool
	switch sortKey {
	case "name":
		switch sortOrder {
		case "asc":
			sortFunc = func(a, b int) bool {
				return strings.Compare(langs[a].name, langs[b].name) < 0
			}
		case "desc":
			sortFunc = func(a, b int) bool {
				return strings.Compare(langs[a].name, langs[b].name) > 0
			}
		default:
			return fmt.Errorf("ListLanguages(): unknown sort order %q", sortOrder)
		}
	case "size":
		switch sortOrder {
		case "asc":
			sortFunc = func(a, b int) bool {
				return langs[a].size < langs[b].size
			}
		case "desc":
			sortFunc = func(a, b int) bool {
				return langs[a].size > langs[b].size
			}
		default:
			return fmt.Errorf("ListLanguages(): unknown sort order %q", sortOrder)
		}
	case "":
		sortFunc = nil
	default:
		return fmt.Errorf("ListLanguages(): unknown sort key %q", sortKey)
	}

	if sortFunc != nil {
		sort.Slice(langs, sortFunc)
	}

	totalSize := 0
	for _, lang := range langs {
		totalSize += lang.size
	}

	totalSizeLabel := "Total size"
	totalSizeString, err := GetSizeByUnit(totalSize, unit)
	if err != nil {
		return err
	}
	maxNameLen := len(totalSizeLabel)
	maxSizeLen := len(totalSizeString)
	for _, lang := range langs {
		if len(lang.name) > maxNameLen {
			maxNameLen = len(lang.name)
		}

		langSize, err := GetSizeByUnit(lang.size, unit)
		if err != nil {
			return err
		}
		langSizeLen := len(langSize)
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
	fmt.Printf("|%s|%s|100.00%%|\n", totalSizeLabel, Strlpad(totalSizeString, maxSizeLen))
	fmt.Println(dashesStr)

	for _, lang := range langs {
		relativeSize := float64(lang.size) / float64(totalSize) * 100
		sizeWithUnit, err := GetSizeByUnit(lang.size, unit)
		if err != nil {
			return err
		}

		fmt.Printf(
			"|%s|%s|%s%%|\n",
			Strrpad(lang.name, maxNameLen),
			Strlpad(sizeWithUnit, maxSizeLen),
			Strlpad(fmt.Sprintf("%.2f", relativeSize), len("100.00")),
		)
	}

	fmt.Println(dashesStr)

	return nil
}
