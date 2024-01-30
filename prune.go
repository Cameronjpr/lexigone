package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func prune(path, key string) {
	f, err := os.Open(path)

	if err != nil {
		fmt.Print(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var bs []byte
	buf := bytes.NewBuffer(bs)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		fmt.Println(key)

		if !strings.Contains(line, key) {
			_, err := buf.WriteString(line + "\n")

			if err != nil {
				fmt.Print(err)
			}
		}
	}

	f.Truncate(0)
	f.Seek(0, 0)
	buf.WriteTo(f)
}
