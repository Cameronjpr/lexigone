package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

func shouldSearchFile(filename string) bool {
	return strings.HasSuffix(filename, ".jsx") || strings.HasSuffix(filename, ".tsx") || strings.HasSuffix(filename, ".js") || strings.HasSuffix(filename, ".ts")
}

func cleanPath(path string) string {
	b1 := strings.ReplaceAll(path, "[", "\\[")
	b2 := strings.ReplaceAll(b1, "]", "\\]")
	b3 := strings.ReplaceAll(b2, "(", "\\(")
	b4 := strings.ReplaceAll(b3, ")", "\\)")
	return b4
}

func search(dir, key string) bool {
	fileSystem := os.DirFS(dir)
	found := false

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, e error) error {
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
			if strings.Contains(line, key) {

				found = true
				return io.EOF
			}
		}

		return nil
	})

	return found
}
