package internal

func EvenSubArray(nums []int32, k int32) int32 {
	// [1,2,3,4] , k=1
	nLen := len(nums)
	start := 0
	var count int32 = 0
	for start < nLen {
		var oddCount int32 = 0
		for end := start; end < nLen; end++ {
			if nums[end]%2 != 0 {
				oddCount++
			}
			if end == nLen-1 {
				start++
			}
			if oddCount > k {
				start++
				break
			} else {
				count++
			}
		}
	}

	return count
}
