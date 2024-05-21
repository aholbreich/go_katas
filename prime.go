package go_kata

import "math"

func IsPrime(n int) bool {
	// Handle edge cases
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	// Any number divisible by 2 or 3 (other than 2 and 3 themselves) is not prime.
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check for factors from 5 to sqrt(n)
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
