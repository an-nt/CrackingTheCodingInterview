package arraysandstrings

// IsUniqueAdvance uses no additional data structure
func IsUniqueAdvance(str string) bool {
	if len(str) > 128 {
		return false
	}
	byteSlice := []byte(str)

	for i, firstCharacter := range byteSlice {
		for _, secondCharacter := range byteSlice[i+1:] {
			if firstCharacter == secondCharacter {
				return false
			}
		}
	}
	return true
}
