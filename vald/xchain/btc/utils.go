package btc

import "math"

func maxInt64(arr []int64) int64 {
	if len(arr) == 0 {
		return math.MinInt64 // Return the smallest possible int64 value if the array is empty
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
