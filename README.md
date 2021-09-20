# lightSearch

## What is this?
lightSearch is a lightweight full-text search engine built for **Markdown** developed in Go.

## Install
-------

    go get github.com/agnivade/levenshtein

## Usage
To use lightSearch, import it and:
```
func main() {
    idx := load("./")
}
```
And search:
```
idx.search("gallery")
```
