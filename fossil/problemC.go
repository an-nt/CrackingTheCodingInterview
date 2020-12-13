package fossil

func ProblemC(n int) string {
	turn := 0
	recorder := make([]string, n)
	return predictWinner(n, turn, recorder)
}
func predictWinner(n int, turn int, recorder []string) string {
	if n < 4 {
		return result[(turn)%2]
	}
	if recorder[n-1] == "" {
		optA := predictWinner(n-2, turn+1, recorder)
		optB := predictWinner(n-3, turn+1, recorder)
		if optA == result[(turn)%2] || optB == result[(turn)%2] {
			recorder[n-1] = result[(turn)%2]
			return recorder[n-1]
		}
		recorder[n-1] = result[(turn+1)%2]
		return recorder[n-1]
	}
	return recorder[n-1]
}

var result = []string{"Bob", "Alice"}
