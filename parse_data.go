package main

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"
)

// ParseData parses a file
// and returns the collections of strings
func ParseData() ([]string, error) {
	var words []string
	fp := path.Join("data", "words.txt")

	file, err := os.Open(fp)
	defer file.Close()
	if err != nil {
		return words, err
	}

	reader := bufio.NewReader(file)
	var line string

	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return words, err
		}

		line = strings.TrimSpace(line)
		if len(line) < 1 {
			continue
		}

		words = append(words, strings.Title(line))
	}

	return words, nil
}
