package arraysandstrings

//IsDuplicated determines whether an interger array contains duplicated values
func IsDuplicated(Arr []int) bool {
	if len(Arr) < 2 {
		return false
	}

	existingValues := make(map[int]bool)

	for _, value := range Arr {
		_, existed := existingValues[value]
		if existed {
			return true
		}
		existingValues[value] = true
	}
	return false
}
