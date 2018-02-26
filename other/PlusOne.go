package other

/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.
*/

func PlusOne(digits []int) []int {

	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
