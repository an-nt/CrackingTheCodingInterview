package SortingAndSearching

func SelectionSort(Arr []int) {
	if len(Arr) < 2 {
		return
	}
	for i := 0; i < len(Arr)-1; i++ {
		minindex := i + 1
		minvalue := Arr[minindex]

		for j := i + 1; j < len(Arr); j++ {
			if minvalue > Arr[j] {
				minindex = j
				minvalue = Arr[minindex]
			}
		}
		if Arr[i] > Arr[minindex] {
			Arr[i], Arr[minindex] = Arr[minindex], Arr[i]
		}
	}
}
