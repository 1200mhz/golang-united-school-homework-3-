package homework

type int64Slice []int64

func reverse(input []int64) (result []int64) {
	//var result []int64

	for i := 0; i < len(input); i++ {
		result = append(result, input[len(input)-i-1])
	}

	return
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
