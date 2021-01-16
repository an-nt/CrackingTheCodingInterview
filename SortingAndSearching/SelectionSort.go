package sortingandsearching

func SelectionSort(Arr []int) {
	if len(Arr) < 2 {
		return
	}
	for i := 0; i < len(Arr)-1; i++ {
		minindex := i
		minvalue := Arr[minindex]

		for j := i; j < len(Arr); j++ {
			if minvalue > Arr[j] {
				minindex = j
				minvalue = Arr[minindex]
			}
		}
		Arr[i], Arr[minindex] = Arr[minindex], Arr[i]
	}
}
