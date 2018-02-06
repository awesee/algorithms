/*

牛年求牛：
	有一母牛，到4岁可生育，每年一头，所生均是一样的母牛
		15岁绝育，不再能生，
		20岁死亡，问n年后有多少头牛。

*/

package main

func main() {

	println(GetCattle(3)) // 1

	println(GetCattle(4)) // 2

	println(GetCattle(10)) // 15

	println(GetCattle(20)) // 346

	println(GetCattle(30)) // 8317
}

var count uint = 1

func GetCattle(n uint) uint {

	var i uint
	for i = 1; i <= n; i++ {
		if i >= 4 && i < 15 {
			count++
			GetCattle(n - i)
		}

		if i == 20 {
			count--
		}
	}

	return count
}
