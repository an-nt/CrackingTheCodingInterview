package arraysandstrings

//IsUnique determines whether a string contains duplicate characters
func IsUnique(str string) bool {
	// ASCII consists of 128 characters
	if len(str) > 128 {
		return false
	}
	characterSet := make(map[byte]bool, 128)

	for _, character := range []byte(str) {
		if characterSet[character] {
			return false
		}
		characterSet[character] = true
	}
	return true
}
