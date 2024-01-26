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
	unused := make([]string, 0)

	for k, v := range total {
		wg.Add(1)
		go func(k, v string) {
			if !search(os.Args[2], k) {
				unused = append(unused, k)
				fmt.Printf("\033[0;31m%s\033[0m in %s\n", k, v)
			}
			wg.Done()
		}(k, v)
	}

	wg.Wait()

	fmt.Printf("Found \033[0;31m%d\033[0m unused keys.\n", len(unused))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
