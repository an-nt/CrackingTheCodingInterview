package main

import (
	"fmt"
)

func main() {
	fmt.Println("Cracking the coding interview")
	// n, q := codesignal.GenerateCase(100000, 1, 300000)

	// start1 := time.Now()
	// fmt.Println(codesignal.SumInRange(n, q))
	// fmt.Println(time.Since(start1))

	// start2 := time.Now()
	// fmt.Println(codesignal.SumInRange_Prefix(n, q))
	// fmt.Println(time.Since(start2))

	// start3 := time.Now()
	// fmt.Println(codesignal.SumInRange_Prefix_Advance(n, q))
	// fmt.Println(time.Since(start3))

	// start4 := time.Now()
	// fmt.Println(codesignal.SumInRange_Prefix_Super(n, q))
	// fmt.Println(time.Since(start4))
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr[6:])
	fmt.Println(arr[:6])
}
