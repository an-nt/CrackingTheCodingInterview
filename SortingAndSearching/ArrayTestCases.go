package sortingandsearching

var ArrayTestCases = [][]int{
	//nil array
	{},
	//arrays with one element
	{-1},
	{1},
	//arrays with two elements
	{1, 2},
	{2, 1},
	{0, 0},
	{-1, -2},
	//arrays with three elements
	{1, 2, 3},
	{1, 3, 2},
	{3, 2, 1},
	{1, 1, 2},
	{1, 2, 1},
	{0, 0, 0},
	{-3, -2, -1},
	{-1, -1, -2},
	{-1, -2, -1},
	//long arrays
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 7, 6, 5, 4, 3, 2, 1},
	{1, 2, 3, 4, 5, 4, 3, 2, 1},
	{9, 8, 7, 6, 5, 6, 7, 8, 9},
	{1, 1, 2, 2, 3, 3, 4, 4, 5},
	{9, 9, 8, 8, 7, 7, 6, 6, 5},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{-4, -3, -2, -1, 0, 1, 2, 3, 4},
	{10, 5, 8, 34, 57, 12, 86, 100, 7, -1, 0, 12, 7, 3, 54, 37},
}
