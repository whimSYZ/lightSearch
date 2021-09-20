package main

func pipline(text string) []string {
	tokens := tokenize(text)
	tokens = stopwordFilter(tokens)
	tokens = lowercaseFilter(tokens)
	tokens = stemmer(tokens)

	return tokens
}
