package sortingandsearching

//MergeSort divides an array into two halves until they become one-element arrays
//merger then merges two corresponding sub-arrays into a sorted array
func MergeSort(Arr []int) []int {
	if len(Arr) < 2 { //condition to stop recursive function
		return Arr
	}
	mid := len(Arr) / 2
	return merger(MergeSort(Arr[:mid]), MergeSort(Arr[mid:])) //divide and merge
}

func merger(firstArr, secondArr []int) []int {
	var i int = 0
	var j int = 0
	var result []int //the array contains sorted elements

	//compare each element in the first array with that in the second
	//the lower one then becomes the new element in the sorted array
	//the comparision stops when all elements of either arrays are added into the result
	for i < len(firstArr) && j < len(secondArr) {
		if firstArr[i] <= secondArr[j] {
			result = append(result, firstArr[i])
			i++
		} else {
			result = append(result, secondArr[j])
			j++
		}
	}

	//the remaining elements are now added into the result
	if i == len(firstArr) {
		result = append(result, secondArr[j:]...)
	} else {
		result = append(result, firstArr[i:]...)
	}
	return result
}
