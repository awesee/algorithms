package other

import "strings"

/*
Write a function to find the longest common prefix string amongst an array of strings.
*/

func LongestCommonPrefix(strs []string) string {

	l := len(strs)
	if l == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < l; i++ {
		for strings.Index(strs[i], pre) != 0 {
			pre = pre[:len(pre)-1]
		}
	}

	return pre
}
