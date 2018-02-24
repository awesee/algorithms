package other

func RemoveDuplicates(nums []int) int {

	n := 1
	for i, v := range nums {
		if i > 0 && nums[n-1] != v {
			nums[n] = v
			n++
		}
	}

	return n
}
