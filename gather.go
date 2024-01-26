package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func getAllKeys(dir string) map[string]string {
	fileSystem := os.DirFS(dir)
	keys := make(map[string]string)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, e error) error {
		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		fileBytes, _ := fs.ReadFile(fileSystem, path)

		if len(fileBytes) == 0 {
			return nil
		}

		var data interface{}
		err := json.Unmarshal(fileBytes, &data)
		check(err)

		m := data.(map[string]interface{})

		for k := range m {
			keys[k] = path
		}

		return nil
	})

	fmt.Printf("Found %d keys\n", len(keys))
	return keys
}
