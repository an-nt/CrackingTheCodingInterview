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
		compressed = append(compressed, inputslice[i])
		start = i
		for j := i; j < len(inputslice); j++ {
			if inputslice[i] != inputslice[j] {
				break
			}
			stop = j
		}

		compressed = append(compressed, []byte(fmt.Sprint(stop-start+1))...)
		if stop > start {
			i = stop
		}
	}

	if len(inputslice) > len(compressed) {
		return string(compressed)
	}
	return string(inputslice)
}
