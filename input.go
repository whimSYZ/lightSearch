package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if len(data) <= 3 {
		//Has no front matter
		return 0, nil, nil
	}

	delim := data[:3]

	if nextDelim := bytes.Index(data[3:], delim); nextDelim > 0 {
		return nextDelim + 6, bytes.TrimSpace(data[:nextDelim+6]), nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}

func read(id int, path string) *markdown {
	var md markdown

	f, err := os.Open(path)

	if err != nil {
		return &md
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(split)

	scanner.Scan()

	frontmatter := strings.TrimSpace(scanner.Text())

	scanner.Scan()

	content := scanner.Text()

	front := make(map[string]interface{})

	yaml.Unmarshal([]byte(frontmatter), front)

	md.id = id
	md.front = front
	md.content = content

	return &md
}

func load(path string) *index {
	var files []string
	ferr := filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if !info.IsDir() && filepath.Ext(p) == ".md" {
			files = append(files, p)
		}
		return nil
	})

	if ferr != nil {
		return nil
	}

	var mds []*markdown

	for i, f := range files {
		mds = append(mds, read(i, f))
	}

	var idx index

	idx.addc(mds)

	return &idx
}
