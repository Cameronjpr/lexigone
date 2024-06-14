package main

import (
	"flag"
	"fmt"
)

func main() {
	var langDirFlag = flag.String("lang", "", "The directory containing the language files")
	var searchDirFlag = flag.String("search", "", "The directory to search for unused keys")
	var cFlag = flag.Bool("c", false, "If -c is set, unused keys are deleted")
	var cleanFlag = flag.Bool("clean", false, "If --clean is set, unused keys are deleted")

	flag.Parse()

	if *langDirFlag == "" {
		panic("Please provide a lang directory with `-lang`")
	}

	if *searchDirFlag == "" {
		panic("Please provide a directory to search with `-search`")
	}

	total := getAllKeys(*langDirFlag)
	unused := findUnused(*searchDirFlag, total)

	if *cFlag || *cleanFlag {
		for k, v := range unused {
			fmt.Printf("Removing %s\n", k)
			prune(v, k)
		}
	}

}
