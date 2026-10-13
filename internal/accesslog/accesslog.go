package accesslog

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Line struct {
	IP     string
	Method string
	Path   string
	Status int
	Bytes  int
}

type Summary struct {
	Parsed    int
	Skipped   int
	UniqueIPs int
	ByStatus  map[int]int
	TopPaths  []PathCount // up to N, count desc, path asc on ties
}

type PathCount struct {
	Path  string
	Count int
}

// ParseLine parses one line.
// Returns ok=false for blank lines, comments, and malformed lines.
func ParseLine(s string) (Line, bool) {
	rawLine := strings.TrimSpace(s)
	if rawLine == "" || rawLine[0] == '#' {
		return Line{}, false
	}

	line := strings.Split(rawLine, " ")
	if len(line) != 5 || line[1] == "" || !strings.HasPrefix(line[2], "/") {
		return Line{}, false
	}

	status, err := strconv.Atoi(line[3])
	if err != nil {
		return Line{}, false
	}

	bytes, err := strconv.Atoi(line[4])
	if err != nil {
		return Line{}, false
	}

	res := Line{
		IP:     line[0],
		Method: line[1],
		Path:   line[2],
		Status: status,
		Bytes:  bytes,
	}

	return res, true
}

// Summarize parses all lines (newline-separated).
// topN is the maximum length of TopPaths. If topN <= 0, TopPaths is empty.
func Summarize(text string, topN int) Summary {
	var res Summary

	lines := strings.Split(text, "\n")
	parsedMap := make(map[string]int)
	uniqueIPs := make(map[string]struct{})
	res.ByStatus = make(map[int]int)

	for _, line := range lines {
		if parsed, ok := ParseLine(line); ok {
			parsedMap[parsed.Path]++
			res.ByStatus[parsed.Status]++
			uniqueIPs[parsed.IP] = struct{}{}
			res.Parsed++
		} else if line != "" {
			res.Skipped++
		}
	}

	if topN > 0 {
		topPaths := make([]PathCount, 0)
		for path, count := range parsedMap {
			topPaths = append(topPaths, PathCount{
				Path:  path,
				Count: count,
			})
		}
		slices.SortFunc(topPaths, func(a, b PathCount) int {
			return cmp.Or(cmp.Compare(b.Count, a.Count), cmp.Compare(a.Path, b.Path))
		})
		res.TopPaths = topPaths[:min(topN, len(topPaths))]
	}

	res.UniqueIPs = len(uniqueIPs)

	return res
}
