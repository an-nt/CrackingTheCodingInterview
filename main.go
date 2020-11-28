package main

import (
	"CodingInterview/recursion"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Cracking the coding interview")
	fmt.Println(recursion.Fibonacci_Memory(100, make([]int, 101)))

	fmt.Println(time.Since(start))
}
