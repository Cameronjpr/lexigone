package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func shouldSearchFile(path string) bool {
	if strings.Contains(path, "node_modules") {
		return false
	}

	return strings.HasSuffix(path, ".jsx") || strings.HasSuffix(path, ".tsx") || strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".ts")
}

func findUnused(dir string, keys map[string]string) map[string]string {
	fileSystem := os.DirFS(dir)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, e error) error {
		fmt.Println("searching:", path)
		if e != nil {
			log.Fatal(e)
		}

		if !shouldSearchFile(path) {
			return nil
		}

		f, err := fs.ReadFile(fileSystem, path)

		if err != nil {
			fmt.Print(err)
			return err
		}

		for _, line := range strings.Split(string(f), "\n") {
			for k := range keys {
				if strings.Contains(line, k) {
					keys[k] = ""
				}
			}
		}

		return nil
	})

	unused := make(map[string]string)

	for k, v := range keys {

		if v != "" {
			unused[k] = v
		}
	}

	return unused
}
