package main

import (
	"flag"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

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
	unused := sync.Map{}

	for k, v := range total {
		wg.Add(1)
		go func(k, v string) {
			if !search(*searchDirFlag, k) {
				unused.Store(k, v)
				fmt.Printf("\033[0;31m%s\033[0m in %s\n", k, v)
			}
			wg.Done()
		}(k, v)
	}

	wg.Wait()

	if *cFlag || *cleanFlag {
		unused.Range(func(key, value interface{}) bool {
			k, ok1 := key.(string)
			p, ok2 := value.(string)
			if !ok1 || !ok2 {
				// handle the case where key or value are not of type string
				return false
			}
			fmt.Printf("Removing %s\n", k)
			prune(p, k)
			return true
		})
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
