package main

import (
	"flag"
	"fmt"
	"os"
)

// name of the executable
var execName = "ghlangs"

// input parameters
var (
	showHelp  bool
	format    string
	sortKey   string
	sortOrder string
	unit      string
)

// init is used here to define input parameters before the execution starts
func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			"Usage: %s [-format FORMAT] [-sort-by KEY] [-sort-order ORDER]\n",
			execName,
		)
		flag.PrintDefaults()
	}

	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showHelp, "h", false, "show help (shorthand)")
	flag.StringVar(&format, "format", "total", "(detail|total) display format")
	flag.StringVar(&sortKey, "sort-by", "size", "(name|size) sort key for sorting languages")
	flag.StringVar(&sortOrder, "sort-order", "desc", "(asc|desc) sort order for sorting languages")
	flag.StringVar(&unit, "unit", "auto", "(auto|B|kB|MB|GB|TB|PB|EB) unit used for displaying sizes")
}

func validateFlags() error {
	switch format {
	case "detail":
	case "total":
	default:
		return fmt.Errorf("unknown display format %q", format)
	}

	switch sortKey {
	case "name":
	case "size":
	default:
		return fmt.Errorf("unknown sort key %q", sortKey)
	}

	switch sortOrder {
	case "asc":
	case "desc":
	default:
		return fmt.Errorf("unknown sort order %q", sortOrder)
	}

	switch unit {
	case "auto":
	case "B":
	case "kB":
	case "MB":
	case "GB":
	case "TB":
	case "PB":
	case "EB":
	default:
		return fmt.Errorf("unknown unit %q", unit)
	}

	return nil
}

func main() {
	flag.Parse()
	if showHelp {
		flag.Usage()
		return
	}

	err := validateFlags()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error: %s\nTry '%s -help' for more information.\n",
			err, execName,
		)
		return
	}

	client, err := NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	repos, err := getRepos(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	langs := getLanguagesFromRepos(repos)
	switch format {
	case "detail":
		err = ListReposWithLanguages(repos, sortKey, sortOrder, unit)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return
		}
	case "total":
		fmt.Println("All repositories:")
		err = ListLanguages(langs, sortKey, sortOrder, unit)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return
		}
	}
}
