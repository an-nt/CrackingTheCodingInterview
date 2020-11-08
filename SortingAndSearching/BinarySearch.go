package SortingAndSearching

//Assume that the array has been sorted in ascending order
func BinarySearch(Arr []int, value int) (index int) {
	if len(Arr) == 0 {
		return -1
	}
	if len(Arr) == 1 {
		if Arr[0] == value {
			return 0
		}
		return -1
	}

	mid := len(Arr) / 2
	if Arr[mid] < value {
		index = BinarySearch(Arr[mid:], value)
		if index == -1 {
			return -1
		}
		return mid + index
	} else if Arr[mid] > value {
		index = BinarySearch(Arr[:mid], value)
		if index == -1 {
			return -1
		}
		return index
	} else {
		return mid
	}
}
