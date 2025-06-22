package utils

func MakeRange(min, max int) []int32 {
	a := make([]int32, max-min+1)
	for index := range a {
		a[index] = int32(min + index)
	}

	return a
}
