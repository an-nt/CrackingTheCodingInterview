package silicon

func add(n int) int {
	if n <= 0 {
		return 0
	}
	return n * (n + 1) / 2
}
