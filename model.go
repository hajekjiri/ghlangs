package main

type langEntry struct {
	name string
	size int
}

type repoEntry struct {
	nameWithOwner string
	langs         []langEntry
}
