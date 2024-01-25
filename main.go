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

	tally := getKeys(args[0])

	for k := range tally {
		search(args[1], k, tally)
	}

	fmt.Println("Unused keys:")
	for k, v := range tally {
		if v == 0 {
			fmt.Printf("%s\n", k)
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
	cmd := exec.Command("grep", "-r", key, dir)
	_, err := cmd.Output()

	if err != nil {
		return false
	}

	tally[key]++

	return true
}

func getKeys(dir string) (keys map[string]int) {
	files, err := os.ReadDir(dir)
	check(err)

	keys = make(map[string]int)

	for _, file := range files {
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
