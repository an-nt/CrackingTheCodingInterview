package sortingandsearching

//QuickSort sorts an array with the complexity of O(N*logN) for average cases and O(N^2) for the worst
//left and right variables indicate the first and the last element of the unsorted array
func QuickSort(Arr []int, left int, right int) {
	if left > right { //condition to stop recursive functions
		return
	}
	part := partition(Arr, left, right)
	QuickSort(Arr, left, part-1)  //Sort the left part
	QuickSort(Arr, part+1, right) //Sort the right part
}

//partition separates an array into two parts
//the left part contains elements lower than the pivot while the other contains those higher
func partition(Arr []int, left int, right int) int {
	//pivot is chosen as the last element (or whatever you want!)
	pivot := Arr[right]
	//part marks the position at which the array is divided into two parts
	part := left
	//iterate all elements but the pivot to determine which one is lower than pivot
	for i := left; i < right; i++ {
		//if an element is lower than the pivot, bring it to the left part
		if Arr[i] < pivot {
			Arr[i], Arr[part] = Arr[part], Arr[i]
			part++
		}
	}
	//bring the pivot to the position between two parts
	Arr[part], Arr[right] = Arr[right], Arr[part]
	return part
}
