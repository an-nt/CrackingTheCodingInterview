package arraysandstrings

func IsUnique(str string) bool {
	// ASCII consists of 128 characters
	if len(str) > 128 {
		return false
	}
	characterSet := make(map[int]bool, 128)

	for _, character := range []byte(str) {
		if characterSet[int(character)] {
			return false
		}
		characterSet[int(character)] = true
	}
	return true
}
