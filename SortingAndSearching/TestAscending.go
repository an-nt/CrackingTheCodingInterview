package SortingAndSearching

func IsAscendOrder(Arr []int) bool {
	if len(Arr) < 2 {
		return true
	}
	for i := 0; i < len(Arr)-1; i++ {
		if Arr[i] > Arr[i+1] {
			return false
		}
	}
	return true
}
