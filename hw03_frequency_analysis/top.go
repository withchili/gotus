package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type pair struct {
	word string
	freq int
}

func Top10(text string) (result []string) {
	words := strings.Fields(text)

	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	pairs := make([]pair, 0, len(freq))
	for w, f := range freq {
		pairs = append(pairs, pair{w, f})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].freq != pairs[j].freq {
			return pairs[i].freq > pairs[j].freq
		}
		return pairs[i].word < pairs[j].word
	})

	n := 10
	if len(pairs) < 10 {
		n = len(pairs)
	}

	for i := 0; i < n; i++ {
		result = append(result, pairs[i].word)
	}

	return result
}
