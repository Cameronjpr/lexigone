package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please provide a lang directory")
	}

	if len(args) < 2 {
		panic("Please provide a directory to search")
	}

	total := getAllKeys(args[0])

	fmt.Println("Total keys:", len(total))

	jobs := make(chan string, len(total))
	results := make(chan string, len(total))

	for w := 1; w <= 3; w++ {
		go worker(jobs, results)
	}

	for k := range total {
		jobs <- total[k]
		fmt.Println("Sent job", total[k])
	}
	close(jobs)

	for i := 0; i < len(total); i++ {
		k := <-results

		if k == "" {
			continue
		}

		fmt.Printf("Key \033[0;31m%s\033[0m is unused.\n", k)
	}

}

func worker(jobs <-chan string, results chan<- string) {
	fmt.Println("Worker started")
	for k := range jobs {
		if search(os.Args[2], k) {
			results <- ""
			continue
		}

		results <- k
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
