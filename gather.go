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

		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", path, err)
			return nil
		}

		m := data.(map[string]interface{})

		for k := range m {
			keys[k] = dir + "/" + path
		}

		return nil
	})

	fmt.Printf("Searching for %d keys\n", len(keys))
	return keys
}
