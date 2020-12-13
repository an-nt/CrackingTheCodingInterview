package codesignal

//FirstDuplicate returns the index of the first duplicated number in an array
func FirstDuplicate(a []int) int {
	recorder := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		_, existed := recorder[a[i]]
		if !existed {
			recorder[a[i]] = true
		} else {
			return a[i]
		}
	}

	return -1
}
