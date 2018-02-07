/*

题目：古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，
	小兔子长到第三个月后每个月又生一对兔子，
	假如兔子都不死，问n个月后的兔子总数为多少？

程序分析：兔子的规律为数列1,1,2,3,5,8,13,21...

*/

package other

func main() {

	println(Rabbit(3)) // 2

	println(Rabbit(10)) // 55

	println(Rabbit(20)) // 6765
}

func Rabbit(n uint) uint {

	if n <= 2 {
		return 1
	}

	return Rabbit(n-1) + Rabbit(n-2)
}
