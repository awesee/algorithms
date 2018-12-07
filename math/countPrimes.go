package math

/*
204. Count Primes

Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

func CountPrimes(n int) int {

	if n < 3 {
		return 0
	}
	primes := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if !primes[i] {
			count++
			for j := i * i; j < n; j += i + i {
				primes[j] = true
			}
		}
	}

	return count
}
