package sortingandsearching

func InsertionSort(Arr []int) {
	if len(Arr) < 2 {
		return
	}

	for i := 1; i < len(Arr); i++ {
		for j := i; j > 0; j-- {
			if Arr[j-1] > Arr[j] {
				Arr[j-1], Arr[j] = Arr[j], Arr[j-1]
			}
		}
	}
}
