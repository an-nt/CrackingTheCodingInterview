package arraysandstrings

func IsPermutation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	if SortString(str1) != SortString(str2) {
		return false
	}
	return true
}

//SortString uses bubble sort algorithm
func SortString(str string) string {
	byteSlice := []byte(str)
	for i := 0; i < len(byteSlice)-1; i++ {
		for j := 0; j < len(byteSlice)-1-i; j++ {
			if int(byteSlice[j]) > int(byteSlice[j+1]) {
				byteSlice[j], byteSlice[j+1] = byteSlice[j+1], byteSlice[j]
			}
		}
	}
	return string(byteSlice)
}
