package codesignal

//FirstNotRepeatingCharacter returns the first unique character in a string
func FirstNotRepeatingCharacter(s string) string {
	inputslice := []byte(s)
	nonrepeatedChar := make(map[byte]int)
	repeatedChar := make(map[byte]int)

	for i := 0; i < len(inputslice); i++ {
		_, existed := nonrepeatedChar[inputslice[i]]
		_, duplicated := repeatedChar[inputslice[i]]
		if !existed && !duplicated {
			nonrepeatedChar[inputslice[i]] = i
		} else {
			if existed {
				repeatedChar[inputslice[i]] = i
				delete(nonrepeatedChar, inputslice[i])
			}
		}
	}
	if len(nonrepeatedChar) == 0 {
		return "_"
	}

	minIndex := len(inputslice)
	for _, index := range nonrepeatedChar {
		if index < minIndex {
			minIndex = index
		}
	}
	return string(inputslice[minIndex])
}
