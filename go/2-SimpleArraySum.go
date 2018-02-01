/*
Given an array of N integers, can you find the sum of its elements?

Input Format

The first line contains an integer, N, denoting the size of the array.
The second line contains N space-separated integers representing the array's elements.

Output Format

Print the sum of the array's elements as a single integer.

Sample Input

6
1 2 3 4 10 11
Sample Output

31
Explanation

We print the sum of the array's elements, which is: 1 + 2 + 3 + 4 + 10 + 11 = 31.
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n uint32
	fmt.Scanf("%v\n", &n)
	var i uint32
	var input = make([]int64, n)
	var sum int64

	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &input[i])
		sum += input[i]
	}

	fmt.Println(sum)
}
