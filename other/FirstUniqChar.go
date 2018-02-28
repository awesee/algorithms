package other

/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

func FirstUniqChar(s string) int {

	a := [256]rune{}
	for _, c := range s {
		a[c]++
	}

	for i, c := range s {
		if a[c] == 1 {
			return i
		}
	}

	return -1
}
