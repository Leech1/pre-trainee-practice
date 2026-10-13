package wordcount

import (
	"slices"
	"strings"
	"unicode"
)

type Pair struct {
	Word  string
	Count int
}

// Tokens splits s into words using the tokenizer rules below.
func Tokens(s string) []string {
	var builder strings.Builder
	res := make([]string, 0)

	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			builder.WriteRune(r)
		} else if builder.Len() > 0 {
			token := strings.ToLower(builder.String())
			res = append(res, token)
			builder.Reset()
		}
	}

	if builder.Len() > 0 {
		res = append(res, strings.ToLower(builder.String()))
	}
	return res
}

// Count returns frequency of each token.
func Count(tokens []string) map[string]int {
	res := make(map[string]int)

	for _, token := range tokens {
		res[token]++
	}
	return res
}

// TopN returns up to n pairs, highest count first.
// Ties: alphabetical ascending by Word.
// If n <= 0, return an empty slice.
func TopN(freq map[string]int, n int) []Pair {
	if n <= 0 || len(freq) == 0 {
		return []Pair{}
	}

	pairs := make([]Pair, 0, len(freq))

	for k, v := range freq {
		pairs = append(pairs, Pair{
			Word:  k,
			Count: v,
		})
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		if a.Count > b.Count {
			return -1
		} else if a.Count < b.Count {
			return 1
		}

		return strings.Compare(a.Word, b.Word)
	})

	return pairs[:min(n, len(pairs))]
}

// DropStop removes tokens that appear in stop (lowercase keys).
func DropStop(tokens []string, stop map[string]struct{}) []string {
	res := make([]string, 0)
	for _, token := range tokens {
		if _, ok := stop[token]; !ok {
			res = append(res, token)
		}
	}
	return res
}
