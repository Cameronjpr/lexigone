package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please provide a lang directory")
	}

	if len(args) < 2 {
		panic("Please provide a directory to search")
	}

	total := getAllKeys(args[0])
	unused := make(map[string]string)

	for k, v := range total {
		wg.Add(1)
		go func(k, v string) {
			if !search(os.Args[2], k) {
				unused[k] = v
				fmt.Printf("\033[0;31m%s\033[0m in %s\n", k, v)
			}
			wg.Done()
		}(k, v)
	}

	wg.Wait()

	for k, p := range unused {
		fmt.Printf("Removing %s\n", k)
		prune(p, k)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
