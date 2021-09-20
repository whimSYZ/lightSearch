# lightSearch
[![https://pkg.go.dev/github.com/whimSYZ/lightSearch](https://pkg.go.dev/badge/github.com/whimSYZ/lightSearch.svg)](https://pkg.go.dev/github.com/whimSYZ/lightSearch)
[![Go Report Card](https://goreportcard.com/badge/github.com/whimSYZ/lightSearch)](https://goreportcard.com/report/github.com/whimSYZ/lightSearch)

Full-text search in Go

## Introduction
lightSearch is a lightweight full-text search engine built for **Markdown** developed in Go. It has optimizations specifically for Markdown documentation such as interpreting Front matter.

## Install

    go get github.com/whimSYZ/lightSearch

## Usage
To use lightSearch:
```go
package main

import (
    "fmt"
    "github.com/whimSYZ/lightSearch"
)

func main() {
    idx := load("./")

    res := idx.search("gallery")

    fmt.Println(res)
}
```

## Dependencies
lightSearch uses [gopkg.in/yaml.v2](gopkg.in/yaml.v2) for YAML Unmarshaler

## Todos
- Add weight calculations for front matter(title, description, etc.)
- Add fuzzy query with levenshteinDistance
