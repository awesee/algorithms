package other

/*
Given an array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
*/

func SingleNumber(nums []int) int {

	n := nums[0]
	for i, v := range nums {
		if i > 0 {
			n ^= v
		}
	}

	return n
}
