package main

type langEntry struct {
	name  string
	bytes int
}

type repoEntry struct {
	nameWithOwner string
	langs         []langEntry
}
