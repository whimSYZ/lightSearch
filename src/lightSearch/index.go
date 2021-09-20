package lightSearch

import (
	"fmt"
	"math"
	"sort"
)

type mdInfo struct{
	id			int
	length 		int
	positions	map[string][]int
}

type index struct {
	fileinf 	map[int]mdInfo //maps id to a map of tokens and its positions in the file with that id
	ind			map[string]token
}

func (idx index) recalculate() {
	for _, token := range idx.ind {
		t := token
		t.indicies = make(map[int]pair)
		for id, file := range idx.fileinf {
			tf := float64( len(file.positions)) / float64(file.length)
			df := t.docCount
			idf := math.Log( float64(len(idx.fileinf)) / float64(df+1))

			p := t.indicies[id]

			p.id = id
			p.val = tf*idf

			t.indicies[id] = p
        }
    }
}

func (idx index) add(files []*markdown) {
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
		f.length = len(processed)
		f.positions = make(map[string][]int)
		
		found := make(map[string]bool)

		for i, token := range processed {
			f.positions[token] = append(f.positions[token], i)
			
			t := idx.ind[token]

			if !found[token] {
				t.docCount++
				found[token] = true
			}
		}

		fmt.Println(f)

		idx.fileinf[file.id] = f
	}
}

func (idx index) addc(files []*markdown) {
	idx.add(files)
	idx.recalculate()
}

func (idx index) search(text string) [][]pair {
    var res [][]pair
    for _, token := range pipline(text) {
		res := []pair{}
		for _, v := range idx.ind[token].indicies {
			res = append(res, v)
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].val > res[j].val
		})
    }
    return res
}