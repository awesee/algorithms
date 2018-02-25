package other

/*
Given an array of integers,
find if the array contains any duplicates.
Your function should return true
if any value appears at least twice in the array,
and it should return false if every element is distinct.
*/

func ContainsDuplicate(nums []int) bool {

	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = true
		}
	}

	return false
}
