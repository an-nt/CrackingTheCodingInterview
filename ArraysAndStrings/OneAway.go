package arraysandstrings

func OneAway(shortString string, longString string) bool {
	lengthDiff := len(longString) - len(shortString)
	if lengthDiff != 0 && lengthDiff != 1 && lengthDiff != -1 {
		return false
	}
	if lengthDiff == -1 {
		return OneAway(longString, shortString)
	}

	if lengthDiff == 0 {
		unmatchCases := 0
		for i := 0; i < len(shortString); i++ {
			if shortString[i] != longString[i] {
				unmatchCases++
			}
			if unmatchCases > 1 {
				return false
			}
		}
		return true
	}

	missCases := 0
	shortIndex := 0
	longIndex := 0
	for shortIndex < len(shortString) {
		if shortString[shortIndex] != longString[longIndex] {
			longIndex++
			missCases++
		}
		if missCases > 1 {
			return false
		}
		shortIndex++
		longIndex++
	}
	return true
}
