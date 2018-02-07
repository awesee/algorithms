/*
Alice and Bob each created one problem for HackerRank.
A reviewer rates the two challenges,
awarding points on a scale from 1 to 100 for three categories:
problem clarity, originality, and difficulty.

We define the rating for Alice's challenge to be the triplet A = (a0, a1, a2),
and the rating for Bob's challenge to be the triplet B = (b0, b1, b2).

Your task is to find their comparison points by comparing  with ,  with , and  with .

If a[i] > b[i], then Alice is awarded  point.
If a[i] < b[i], then Bob is awarded  point.
If a[i] = b[i], then neither person receives a point.
Comparison points is the total points a person earned.

Given A and B, can you compare the two challenges and print their respective comparison points?

Input Format

The first line contains 3 space-separated integers, a0, a1, and a2, describing the respective values in triplet .
The second line contains 3 space-separated integers, b0, b1, and b2, describing the respective values in triplet .

Constraints
  1<=a[i]<=100
  1<=b[i]<=100


Output Format

Print two space-separated integers denoting the respective comparison points earned by Alice and Bob.

Sample Input

5 6 7
3 6 10
Sample Output

1 1
Explanation

In this example:
A = (a0, a1, a2) = (5, 6, 7)
B = (b0, b1, b2) = (3, 6, 10)

Now, let's compare each individual score:

a[0] > b[0], so Alice receives 1 point.
a[1] = b[1], so nobody receives a point.
a[2] < b[2], so Bob receives 1 point.
Alice's comparison score is 1, and Bob's comparison score is 1. Thus, we print 1 1 (Alice's comparison score followed by Bob's comparison score) on a single line.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		A = make([]string, 0)
		B = make([]string, 0)
	)

	input := bufio.NewScanner(os.Stdin)
	end := 0

	for input.Scan() {
		t := input.Text()
		if end == 0 {
			A = strings.Split(t, " ")
		} else {
			B = strings.Split(t, " ")
			break
		}

		end++
	}

	var a, b int64
	for i := range A {
		x, _ := strconv.Atoi(A[i])
		y, _ := strconv.Atoi(B[i])
		if x > y {
			a++
		} else if x < y {
			b++
		}
	}

	fmt.Println(a, b)
}
