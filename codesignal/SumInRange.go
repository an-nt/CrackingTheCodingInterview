package codesignal

import "math/rand"

func GenerateCase(numsLen int, value int, queriesLen int) ([]int, [][]int) {
	//generate nums
	nums := make([]int, numsLen)
	for i := 0; i < len(nums); i++ {
		nums[i] = value
	}
	//generate queries
	queries := make([][]int, queriesLen)
	for j := 0; j < len(queries); j++ {
		min := rand.Int() % len(nums)
		max := min + (rand.Int() % (len(nums) - min))
		queries[j] = []int{min, max}
	}
	return nums, queries
}

func SumInRange(nums []int, queries [][]int) int {
	sum := 0
	for _, query := range queries {
		for i := query[0]; i <= query[1]; i++ {
			sum += nums[i]
		}
	}
	return sum
}

func SumInRange_Prefix(nums []int, queries [][]int) int {
	sum := 0
	queries = MergeSortMatrix(queries)
	for i := 0; i < len(queries); {
		start := i
		memory := make(map[int]int)
		for i < len(queries) && queries[start][0] == queries[i][0] {
			if i == start {
				tempSum := 0
				for k := queries[start][0]; k <= queries[start][1]; k++ {
					tempSum += nums[k]
				}
				memory[queries[i][0]+queries[i][1]] = tempSum
				sum += memory[queries[i][0]+queries[i][1]]
			}
			if i > start {
				_, existed := memory[queries[i][0]+queries[i][1]]
				if !existed {
					tempSum := memory[queries[i-1][0]+queries[i-1][1]]
					for k := queries[i-1][1] + 1; k <= queries[i][1]; k++ {
						tempSum += nums[k]
					}
					memory[queries[i][0]+queries[i][1]] = tempSum
				}
				sum += memory[queries[i][0]+queries[i][1]]
			}
			i++
		}
	}
	return sum
}

func SumInRange_Prefix_Advance(nums []int, queries [][]int) int {
	sum := 0
	queries = MergeSortMatrix(queries)
	previousSum := 0
	tempSum := 0
	for i := 0; i < len(queries); {
		memory := make(map[int]int)
		start := i
		if i == 0 {
			tempSum = 0
			for k := queries[0][0]; k <= queries[0][1]; k++ {
				tempSum += nums[k]
			}
			memory[queries[0][0]+queries[0][1]] = tempSum
			previousSum = memory[queries[0][0]+queries[0][1]]
			sum += memory[queries[0][0]+queries[0][1]]
			i++
		}

		for i < len(queries) && queries[start][0] == queries[i][0] {
			if i == start {
				if queries[i-1][1] < queries[i][0] {
					tempSum = 0
					for k := queries[start][0]; k <= queries[start][1]; k++ {
						tempSum += nums[k]
					}
				} else {
					tempSum = previousSum
					for t := queries[i-1][0]; t < queries[i][0]; t++ {
						tempSum -= nums[t]
					}
					if queries[i-1][1] >= queries[i][1] {
						for k := queries[i][1] + 1; k <= queries[i-1][1]; k++ {
							tempSum -= nums[k]
						}
					} else {
						for h := queries[i-1][1] + 1; h <= queries[i][1]; h++ {
							tempSum += nums[h]
						}
					}
				}
				memory[queries[i][0]+queries[i][1]] = tempSum
				previousSum = memory[queries[i][0]+queries[i][1]]
				sum += memory[queries[i][0]+queries[i][1]]
			}
			if i > start {
				_, existed := memory[queries[i][0]+queries[i][1]]
				if !existed {
					tempSum := memory[queries[i-1][0]+queries[i-1][1]]
					for k := queries[i-1][1] + 1; k <= queries[i][1]; k++ {
						tempSum += nums[k]
					}
					memory[queries[i][0]+queries[i][1]] = tempSum
					previousSum = memory[queries[i][0]+queries[i][1]]
				}
				sum += memory[queries[i][0]+queries[i][1]]
			}
			i++
		}
	}
	return sum
}

func MergeSortMatrix(m [][]int) [][]int {
	if len(m) < 2 {
		return m
	}
	mid := len(m) / 2
	return merge(MergeSortMatrix(m[:mid]), MergeSortMatrix(m[mid:]))
}

func merge(m1 [][]int, m2 [][]int) [][]int {
	result := [][]int{}
	i := 0
	j := 0
	for i < len(m1) && j < len(m2) {
		if m1[i][0] < m2[j][0] {
			result = append(result, m1[i])
			i++
		} else if m1[i][0] > m2[j][0] {
			result = append(result, m2[j])
			j++
		} else {
			if m1[i][1] <= m2[j][1] {
				result = append(result, m1[i])
				i++
			} else {
				result = append(result, m2[j])
				j++
			}
		}
	}
	if i < len(m1) {
		return append(result, m1[i:]...)
	}
	return append(result, m2[j:]...)
}

func SumInRange_Prefix_Super(nums []int, queries [][]int) int {
	sum := 0
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for _, query := range queries {
		if query[0] == 0 {
			sum += nums[query[1]]
		} else {
			sum += nums[query[1]] - nums[query[0]-1]
		}
	}
	return sum
}
