package SortingAndSearching

func BubbleSort(Arr []int) {
	if len(Arr) < 2 {
		return
	}

	for i := 0; i < len(Arr); i++ {
		for j := 0; j < len(Arr)-1; j++ {
			if Arr[j] > Arr[j+1] {
				Arr[j], Arr[j+1] = Arr[j+1], Arr[j]
			}
		}
	}
}
