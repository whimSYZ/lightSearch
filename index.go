package main

import (
	"fmt"
	"math"
	"sort"
)

type mdInfo struct {
	id        int
	positions map[string][]int
}

type index struct {
	fileinf map[int]mdInfo //maps id to a map of tokens and its positions in the file with that id
	ind     map[string]token
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func levenshteinDistance(a, b string) int {
	str1 := []rune(a)
	str2 := []rune(b)

	c := make([]int, len(str1)+1)

	for i := 1; i < len(c); i++ {
		c[i] = i
	}

	for i := 1; i <= len(str2); i++ {
		pre := i
		for j := 1; j <= len(str1); j++ {
			cur := c[j-1]
			if str1[j-1] != str2[i-1] {
				cur = min(min(c[j-1]+1, c[j]+1), pre+1)
			}
			c[j-1] = pre
			pre = cur
		}
		c[len(str1)] = pre
	}
	return int(c[len(str1)])
}

func (idx *index) recalculate() {
	for token, t := range idx.ind {
		t.indicies = make(map[int]pair)

		for id, file := range idx.fileinf {
			//fmt.Println(file.positions)
			tf := float64(len(file.positions[token])) / float64(len(file.positions))

			df := t.docCount
			idf := math.Log(float64(len(idx.fileinf)) / float64(df+1))

			p := t.indicies[id]

			p.id = id
			p.val = tf * idf

			t.indicies[id] = p
		}

		idx.ind[token] = t
	}
}

func (idx *index) add(files []*markdown) {
	if idx.fileinf == nil {
		idx.fileinf = make(map[int]mdInfo)
	}
	if idx.ind == nil {
		idx.ind = make(map[string]token)
	}

	for _, file := range files {
		processed := pipline(file.content)
		f := idx.fileinf[file.id]
		f.id = file.id
		f.positions = make(map[string][]int)

		found := make(map[string]bool)

		for i, token := range processed {
			if len(token) > 0 {
				f.positions[token] = append(f.positions[token], i)

				t := idx.ind[token]
				t.str = token

				if !found[token] {
					t.docCount++
					found[token] = true
				}

				idx.ind[token] = t
			}
		}
		idx.fileinf[file.id] = f
	}
}

func (idx *index) addc(files []*markdown) {
	idx.add(files)

	idx.recalculate()

	fmt.Println(idx)
}

func (idx *index) search(text string) []pair {
	r := make(map[int]float64)

	for _, token := range pipline(text) {
		for id, p := range idx.ind[token].indicies {
			r[id] += p.val
		}
	}

	docs := []pair{}

	for id, val := range r {
		if val != 0 {
			docs = append(docs, pair{id, val})
		}
	}

	sort.Slice(docs, func(i, j int) bool {
		return docs[i].val > docs[j].val
	})

	/*
		res := []int{}
		for _, v := range docs {
			res = append(res, v.id)
		}*/

	return docs
}
