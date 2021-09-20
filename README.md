# lightSearch

## What is this?
lightSearch is a lightweight full-text search engine built for **Markdown** developed in Go. It has optimizations specifically for Markdown documentation such as interpreting Front matter.

## Install

    go get github.com/whimSYZ/lightSearch

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

## Dependencies
lightSearch uses 
