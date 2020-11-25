package arraysandstrings

func PalindromePermutation(str string) bool {
	recorder := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		if str[i] != []byte(" ")[0] {
			recorder[str[i]]++
		}
	}

	var oddChar int = 0
	for _, occurrence := range recorder {
		if occurrence%2 == 1 {
			oddChar++
		}
		if oddChar > 1 {
			return false
		}
	}
	return true
}
