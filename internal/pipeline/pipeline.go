package pipeline

// Double returns a new slice where each element is n*2.
// If in == nil, return nil.
func Double(in []int) []int {
	if in == nil {
		return nil
	}

	out := make([]int, 0, len(in))
	for _, num := range in {
		out = append(out, num*2)
	}
	return out
}

// FilterMin returns a new slice with elements >= min, in the same order.
// If in == nil, return nil.
func FilterMin(in []int, min int) []int {
	if in == nil {
		return nil
	}

	out := make([]int, 0)
	for _, num := range in {
		if num >= min {
			out = append(out, num)
		}
	}
	return out
}

// UniqueStable returns first occurrences only, original order kept.
// If in == nil, return nil.
func UniqueStable(in []int) []int {
	if in == nil {
		return nil
	}
	out := make([]int, 0)
	occurrences := make(map[int]struct{})

	for _, num := range in {
		if _, ok := occurrences[num]; !ok {
			out = append(out, num)
			occurrences[num] = struct{}{}
		}
	}
	return out
}
