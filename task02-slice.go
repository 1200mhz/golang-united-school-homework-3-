package homework

import "sort"

type int64Slice []int64

func reverse(input []int64) []int64 {
	res := int64Slice{}
	res = input

	sort.Sort(sort.Reverse(&res))

	return res
}

func (s *int64Slice) Len() int {
	tl := 0
	for range *s {
		tl += 1
	}

	return tl
}

func (s *int64Slice) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *int64Slice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
