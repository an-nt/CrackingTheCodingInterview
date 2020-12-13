package recursion

// CanSum_memo determines whether it is possible to generate a target sum
// using elements in a nonnegative-integer array
// the element can be used any time
func CanSum_memo(targetSum int, nums []int) bool {
	if targetSum < 0 {
		return false
	}

	memo := make(map[int]bool)
	memo[0] = true
	return cansum_memo(targetSum, nums, memo)
}

// Complexity: m = targetSum, n = len(nums)
// time: n*m
// space: m(m+n)

func cansum_memo(targetSum int, nums []int, memo map[int]bool) bool {
	_, existed := memo[targetSum]
	if existed {
		return memo[targetSum]
	}
	if targetSum < 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			subResult := cansum_memo(targetSum-nums[i], nums, memo)
			if subResult {
				memo[targetSum] = true
				return true
			}
		}
	}
	memo[targetSum] = false
	return false
}

// CanSum_tabu is similar CanSum_memo, but use another technique
func CanSum_tabu(targetSum int, nums []int) bool {
	if targetSum < 0 {
		return false
	}
	array := make([]bool, targetSum+1)

	array[0] = true
	for i := 0; i < len(array); i++ {
		if array[i] == true {
			for _, num := range nums {
				if i+num < len(array) {
					array[i+num] = true
				}
			}
		}
	}
	return array[targetSum]
}

// Complexity: m = targetSum, n = len(nums)
// time: m*n
// space: m
