package arraysandstrings

import "fmt"

func StringCompression(str string) string {
	if len(str) < 2 {
		return str
	}
	inputslice := []byte(str)
	compressed := []byte("")
	start := 0
	stop := 0
	for i := 0; i < len(inputslice); i++ {
		// Add the i(th) character to the result
		compressed = append(compressed, inputslice[i])
		start = i
		// Determine whether there are repeated characters
		for j := i; j < len(inputslice); j++ {
			if inputslice[i] != inputslice[j] {
				break
			}
			stop = j
		}
		// Add the number of character appearing in the string
		compressed = append(compressed, []byte(fmt.Sprint(stop-start+1))...)
		// Move i to the last repeated character
		if stop > start {
			i = stop
		}
	}
	// Determine whether the compressed string is shorter than the original
	if len(inputslice) > len(compressed) {
		return string(compressed)
	}
	return string(inputslice)
}
