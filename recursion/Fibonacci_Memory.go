package recursion

func Fibonacci_Memory(n int, memory []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memory[n] == 0 {
		memory[n] = Fibonacci_Memory(n-1, memory) + Fibonacci_Memory(n-2, memory)
	}
	return memory[n]
}
