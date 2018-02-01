/*
Complete the function solveMeFirst to compute the sum of two integers.

Function prototype:

int solveMeFirst(int x, int y);

where,

x is the first integer input.
y is the second integer input
Return values

sum of the above two integers
Sample Input

x = 2
y = 3
Sample Output

5
Explanation

The sum of the two integers x and y is computed as: 2 + 3 = 5.
*/

package main

import "fmt"

func solveMeFirst(a, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
