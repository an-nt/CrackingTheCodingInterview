package fossil

func ProblemA(y int) int {
	numOfDegits := countDegits(y)
	minPossibleAnswer := y - numOfDegits*9
	if minPossibleAnswer < 0 {
		minPossibleAnswer = 1
	}
	for possibleAnswer := minPossibleAnswer; possibleAnswer < y; possibleAnswer++ {
		a := possibleAnswer
		for j := a; j > 0; j /= 10 {
			lastDegit := j % 10
			a += lastDegit
		}

		if y == a {
			return possibleAnswer
		}
	}
	return -1
}

func countDegits(n int) int {
	count := 0
	for i := n; i > 0; i /= 10 {
		count++
	}
	return count
}
