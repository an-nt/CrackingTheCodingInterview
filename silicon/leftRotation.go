package silicon

func leftRotation(array []int, d int) []int {
	if d < 1 || d > len(array) {
		return array
	}
	return append(array[d:], array[:d]...)
}
