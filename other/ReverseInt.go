package other

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func ReverseInt(x int) int {

	res, t := 0, 0
	for x != 0 {
		t = res*10 + x%10
		if t/10 != res {
			return 0
		}
		res = t
		x /= 10
	}

	return res
}
