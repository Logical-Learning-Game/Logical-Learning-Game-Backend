package utility

func SliceContain[T comparable](elements []T, value T) bool {
	for _, v := range elements {
		if v == value {
			return true
		}
	}
	return false
}

func TwoDimensionSlice[T interface{}](oneDimensionSlice []T, height, width int) [][]T {
	result := make([][]T, height)

	for i := 0; i < height; i++ {
		innerSlice := make([]T, width)

		for j := 0; j < width; j++ {
			innerSlice[j] = oneDimensionSlice[width*i+j]
		}

		result[i] = innerSlice
	}

	return result
}

func ReverseSlice[T interface{}](s []T) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
