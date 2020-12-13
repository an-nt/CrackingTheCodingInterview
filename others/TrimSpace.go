package others

func TrimSpace(str string) string {
	//remove the space at the beginning
	for str != "" {
		if string(str[0]) != " " {
			break
		}
		str = str[1:]
	}

	//remove the space at the end
	for str != "" {
		if string(str[len(str)-1]) != " " {
			break
		}
		str = str[:len(str)-1]
	}

	if str != "" {
		for i := 0; i < len(str); i++ {
			if string(str[i]) == " " {
				j := i + 1
				for j < len(str) && string(str[j]) == " " {
					j++
				}
				if j-i > 1 {
					str = string(append([]byte(str[:i+1]), str[j:]...))
				}
			}
		}
	}

	return str
}
