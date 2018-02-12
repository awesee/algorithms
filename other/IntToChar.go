/*

数字(num >= 0)转字符(A-Z)

把数字转为Excel列坐标值

例如：
	0 => A
	1 => B
	...
	25 => Z
	26 => AA
	27 => AB
	...

*/

package other

import (
	"math"
)

func IntToChar(num uint) (char string) {

	if quotient := math.Floor(float64(num / 26)); quotient > 0 {

		char += IntToChar(uint(quotient - 1))
	}

	return char + string(num%26+65)
}
