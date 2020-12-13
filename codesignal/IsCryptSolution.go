package codesignal

import "strconv"

//IsCryptSolution checks whether a crypt and a solution are matching or not
func IsCryptSolution(crypt []string, solution [][]string) bool {
	equation := []int{}
	for i := 0; i < len(crypt); i++ {
		num := stringToNum(crypt[i], solution)
		if num == -1 {
			return false
		}
		equation = append(equation, num)
	}
	if equation[0]+equation[1] != equation[2] {
		return false
	}
	return true
}

func stringToNum(str string, solution [][]string) int {
	inputslice := []byte(str)
	result := 0
	num := 0
	for i := 0; i < len(inputslice); i++ {
		for j := 0; j < len(solution); j++ {
			if solution[j][0] == string(inputslice[i]) {
				num, _ = strconv.Atoi(solution[j][1])
				break
			}
		}
		if i == 0 && num == 0 && len(inputslice) != 1 {
			return -1
		}
		result = result*10 + num
	}
	return result
}
