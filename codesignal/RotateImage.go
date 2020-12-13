package codesignal

//RotateImage rotates an matrix (NxN) by 90 degrees (clockwise)
func RotateImage(a [][]int) [][]int {
	if len(a) < 2 {
		return a
	}
	for i := 0; i < (len(a)+1)/2; i++ {
		for j := 0; j < len(a)/2; j++ {
			a[i][j], a[j][len(a)-1-i], a[len(a)-1-i][len(a)-1-j], a[len(a)-1-j][i] = a[len(a)-1-j][i], a[i][j], a[j][len(a)-1-i], a[len(a)-1-i][len(a)-1-j]
		}
	}
	return a
}
