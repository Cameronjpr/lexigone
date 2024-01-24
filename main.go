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
		fmt.Println("Please provide a directory")
		os.Exit(1)
	}

	tally := getKeys(args[0])

	for k := range tally {
		cmd := exec.Command("grep", "-r", k, args[1])
		_, err := cmd.Output()

		if err != nil {
			continue
		}

		tally[k]++
	}

	unused := make(map[string]int, len(tally))

	for k, v := range tally {
		if v == 0 {
			unused[k] = v
		}
	}

	for k := range unused {
		fmt.Printf("%s\n", k)
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

func getKeys(dir string) (keys map[string]int) {
	files, err := os.ReadDir(dir)
	check(err)

	for _, file := range files {
		if strings.Split(file.Name(), ".")[1] != "json" {
			continue
		}

		fileBytes, _ := os.ReadFile(dir + "/" + file.Name())

		var data interface{}
		err = json.Unmarshal(fileBytes, &data)
		check(err)

		m := data.(map[string]interface{})
		keys = make(map[string]int, len(m))

		for k := range m {
			keys[k] = 0
		}
	}

	return keys
}
