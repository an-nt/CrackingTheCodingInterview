package recursion

func Fibonacci_Recursion(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci_Recursion(n-1) + Fibonacci_Recursion(n-2)
}
