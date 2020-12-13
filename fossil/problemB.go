package fossil

func ProblemB(s string) string {
	countA := 0
	countB := 0
	countC := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			countA++
		} else if string(s[i]) == "B" {
			countB++
		} else {
			countC++
		}
	}

	if countA < len(s)/3 {
		for countA < len(s)/3 && countB > len(s)/3 {
			s = replace(s, "B", "A", true)
			countB--
			countA++
		}
		for countA < len(s)/3 && countC > len(s)/3 {
			s = replace(s, "C", "A", true)
			countC--
			countA++
		}
	}
	if countB < len(s)/3 {
		for countB < len(s)/3 && countA > len(s)/3 {
			s = replace(s, "A", "B", false)
			countB++
			countA--
		}
		for countB < len(s)/3 && countC > len(s)/3 {
			s = replace(s, "C", "B", true)
			countB++
			countC--
		}
	}

	if countC < len(s)/3 {
		for countC < len(s)/3 && countA > len(s)/3 {
			s = replace(s, "A", "C", false)
			countC++
			countA--
		}
		for countC < len(s)/3 && countB > len(s)/3 {
			s = replace(s, "B", "C", false)
			countC++
			countB--
		}
	}
	return s
}

func replace(str string, replace string, by string, replaceFirst bool) string {
	inputslice := []byte(str)
	if replaceFirst {
		for i := 0; i < len(inputslice); i++ {
			if string(inputslice[i]) == replace {
				inputslice[i] = []byte(by)[0]
				break
			}
		}
		return string(inputslice)
	}
	for i := len(str) - 1; i >= 0; i-- {
		if string(inputslice[i]) == replace {
			inputslice[i] = []byte(by)[0]
			break
		}
	}
	return string(inputslice)
}
