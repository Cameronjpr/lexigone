package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

func shouldSearchFile(filename string) bool {
	return strings.HasSuffix(filename, ".jsx") || strings.HasSuffix(filename, ".tsx")
}

func prepareCommand(cmd string) string {
	b1 := strings.ReplaceAll(cmd, "[", "\\[")
	b2 := strings.ReplaceAll(b1, "]", "\\]")
	b3 := strings.ReplaceAll(b2, "(", "\\(")
	b4 := strings.ReplaceAll(b3, ")", "\\)")
	return b4
}

func search(dir, key string) bool {
	fileSystem := os.DirFS(dir)
	found := false

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, e error) error {
		if !shouldSearchFile(path) {
			return nil
		}

		cmd := prepareCommand(fmt.Sprintf("grep %s %s", key, path))

		ex := exec.Command("bash", "-c", cmd)

		out, err := ex.Output()

		fmt.Println(string(out))

		if err == nil {
			fmt.Println("Found", key, "in", path)
			found = true
			return io.EOF
		}

		return nil
	})

	return found
}
