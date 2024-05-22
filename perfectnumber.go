package go_kata

import (
	"math/big"
)

func isPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func FindPerfectNumbers(n int) []int {
	perfectNumbers := []int{}
	num := 1
	for len(perfectNumbers) < n {
		if isPerfectNumber(num) {
			perfectNumbers = append(perfectNumbers, num)
		}
		num++
	}
	return perfectNumbers
}

func FindPerfectNumbersEfficient(n int) []*big.Int {

	var perfectNumbers = make([]*big.Int, n)
	count := 0

	var bigInt = new(big.Int)
	var base = big.NewInt(2)

	for i := 2; count < n; i++ {

		if !big.NewInt(int64(i)).ProbablyPrime(20) {
			continue
		}

		var e = big.NewInt(int64(i))
		var mersennePrime = bigInt.Exp(base, e, nil)
		mersennePrime = bigInt.Sub(mersennePrime, big.NewInt(1))

		if mersennePrime.ProbablyPrime(20) {

			e = big.NewInt(int64(i - 1))
			mersennePrime2 := new(big.Int).Exp(base, e, nil)

			perfectNumbers[count] = new(big.Int).Mul(mersennePrime, mersennePrime2)
			count++
		}
	}

	return perfectNumbers
}
