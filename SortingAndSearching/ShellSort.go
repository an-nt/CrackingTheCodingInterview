package sortingandsearching

func ShellSort(arr []int) {
	//determine a list of user-defined intervals that decrease sequentially
	intervals := intervalsGenerator(arr)

	//use insertion sort to sort a sub-array that is formed by interval-separating elements
	for _, interval := range intervals {
		for i := interval; i < len(arr); i += interval {
			for j := i; j > 0; j -= interval {
				if arr[j-interval] > arr[j] {
					arr[j-interval], arr[j] = arr[j], arr[j-interval]
				}
			}
		}
	}

}

//this is just one posible interval slice
//this function will be modified later
func intervalsGenerator(arr []int) []int {
	return []int{701, 301, 132, 57, 23, 10, 4, 1}
}
