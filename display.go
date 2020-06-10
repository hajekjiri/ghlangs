package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func listReposWithLanguages(repos []repoEntry, unit string) {
	for i, repo := range repos {
		fmt.Printf("%s\n", repo.nameWithOwner)

		listLanguages(repo.langs, "bytes", "descending", unit)

		if i < len(repos)-1 {
			fmt.Println()
		}
	}
}

func listLanguages(langs []langEntry, sortKey string, sortDirection string, unit string) {
	var sortFunc func(a, b int) bool
	switch sortKey {
	case "name":
		switch sortDirection {
		case "ascending":
			sortFunc = func(a, b int) bool {
				return strings.Compare(langs[a].name, langs[b].name) < 0
			}
		case "descending":
			sortFunc = func(a, b int) bool {
				return strings.Compare(langs[a].name, langs[b].name) > 0
			}
		default:
			log.Printf(
				"Warning: unknown sort direction '%s' in listLanguages(), defaulting to 'ascending'\n",
				sortDirection,
			)
			sortFunc = func(a, b int) bool {
				return strings.Compare(langs[a].name, langs[b].name) < 0
			}
		}
	case "bytes":
		switch sortDirection {
		case "ascending":
			sortFunc = func(a, b int) bool {
				return langs[a].bytes < langs[b].bytes
			}
		case "descending":
			sortFunc = func(a, b int) bool {
				return langs[a].bytes > langs[b].bytes
			}
		default:
			log.Printf(
				"Warning: unknown sort direction '%s' in listLanguages(), defaulting to descending\n",
				sortDirection,
			)
			sortFunc = func(a, b int) bool {
				return langs[a].bytes > langs[b].bytes
			}
		}
	case "":
		sortFunc = nil
	default:
		log.Printf("Warning: unknown sort key '%s' in listLanguages()\n", sortKey)
		sortFunc = nil
	}

	if sortFunc != nil {
		sort.Slice(langs, sortFunc)
	}

	totalSize := 0
	for _, lang := range langs {
		totalSize += lang.bytes
	}

	totalSizeLabel := "Total size"
	totalSizeString, err := GetSizeByUnit(totalSize, unit)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	maxNameLen := len(totalSizeLabel)
	maxSizeLen := len(totalSizeString)
	for _, lang := range langs {
		if len(lang.name) > maxNameLen {
			maxNameLen = len(lang.name)
		}

		langSize, err := GetSizeByUnit(lang.bytes, unit)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
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
		relativeSize := float64(lang.bytes) / float64(totalSize) * 100
		size, err := GetSizeByUnit(lang.bytes, unit)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		fmt.Printf(
			"|%s|%s|%s%%|\n",
			Strrpad(lang.name, maxNameLen),
			Strlpad(size, maxSizeLen),
			Strlpad(fmt.Sprintf("%.2f", relativeSize), len("100.00")),
		)
	}
	fmt.Println(dashesStr)
}
