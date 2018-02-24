package other

/*
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.

[show hint]

Related problem: Reverse Words in a String II

Credits:
Special thanks to @Freezen for adding this problem and creating all test cases.
*/

func RotateArray(nums []int, k int) {

	l := len(nums)
	k %= l
	s := make([]int, l)
	copy(s, nums)

	for i, v := range s[l-k:] {
		nums[i] = v
	}

	for i, v := range s[:l-k] {
		nums[k+i] = v
	}
}
