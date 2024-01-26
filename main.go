package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide a lang directory")
		os.Exit(1)
	}

	if len(args) < 2 {
		fmt.Println("Please provide a search directory")
		os.Exit(1)
	}

	keys := make(map[string]int)
	tally := getKeys(args[0], keys)

	for k := range tally {
		fmt.Println("Searching for", k)
		search(args[1], k, tally)
	}

	fmt.Println("")
	fmt.Println("Unused keys:")
	for k, v := range tally {
		if v == 0 {
			fmt.Printf("\033[0;31m%s\n\033[0m", k)
		}
	}
}

type Data struct {
	Keys []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func search(dir string, key string, tally map[string]int) bool {
	cmd := fmt.Sprintf("grep -r --exclude='*.json' %s %s", key, dir)
	ex := exec.Command("bash", "-c", cmd)

	_, err := ex.CombinedOutput()

	if err != nil {
		return false
	}

	fmt.Println("Found", key)

	tally[key]++

	return true
}

func getKeys(dir string, keys map[string]int) map[string]int {
	files, err := os.ReadDir(dir)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Searching", dir+"/"+file.Name())
			getKeys(dir+"/"+file.Name(), keys)
			continue
		}

		fmt.Println("Reading", dir+"/"+file.Name())
		if strings.Split(file.Name(), ".")[1] != "json" {
			continue
		}

		fileBytes, _ := os.ReadFile(dir + "/" + file.Name())

		var data interface{}
		err = json.Unmarshal(fileBytes, &data)
		check(err)

		m := data.(map[string]interface{})

		for k := range m {
			keys[k] = 0
		}
	}

	return keys
}
