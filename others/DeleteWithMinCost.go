package others

func DeleteWithMinCost(str string, cost []int) int {
	var start int
	var stop int
	var minCost int

	if len(str) != len(cost) {
		return minCost
	}

	for i := 0; i < len(str)-1; i++ {
		start = i

		for j := i + 1; j < len(str); j++ {
			if str[i] != str[j] {
				break
			}
			stop = j
		}

		if stop > start {
			minCost += getMinCost(cost[start : stop+1])
			i = stop
		}
	}
	return minCost
}

func getMinCost(arr []int) int {
	var max int
	var sum int

	if len(arr) == 0 {
		return sum
	}
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] > max {
			max = arr[i]
		}
	}
	return sum - max
}
