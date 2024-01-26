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

	for _, v := range total {
		wg.Add(1)
		go func(v string) {
			found := search(os.Args[2], v)
			fmt.Println(v, found)
			if !search(os.Args[2], v) {
				unused = append(unused, v)
			}
			wg.Done()
		}(v)
	}

	wg.Wait()

	for _, v := range unused {
		fmt.Printf("Key \033[0;31m%s\033[0m is unused.\n", v)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
