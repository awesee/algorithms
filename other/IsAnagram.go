package other

/*
Given two strings s and t, write a function to determine if t is an anagram of s.

For example,
s = "anagram", t = "nagaram", return true.
s = "rat", t = "car", return false.

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	as := [256]rune{}
	for _, v := range s {
		as[v]++
	}

	for _, v := range t {
		as[v]--
		if as[v] < 0 {
			return false
		}
	}

	return true
}
