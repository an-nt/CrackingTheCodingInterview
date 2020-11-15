package arraysandstrings

func IsPermutation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	if SortStringBubble(str1) != SortStringBubble(str2) {
		return false
	}
	return true
}

//SortStringBubble uses bubble sort algorithm
func SortStringBubble(str string) string {
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

func SortStringQuickSort(str string) string {
	return ""
}
