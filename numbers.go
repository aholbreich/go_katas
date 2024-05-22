package go_kata

import "math/big"

// shifts the integer n to the left by i bits.
func ShiftLeft(n int, i uint) int {
	return n << i
}

// shifts the big integer n to the left by i bits.
func ShiftLeftBig(n uint, i uint) big.Int {
	var res = big.NewInt(1)
	return *res.Lsh(big.NewInt(int64(n)), uint(i))
}

// shifts the big integer n to the left by 1 bit.
func ShiftLeftMultBy2(n uint) big.Int {
	var res = big.NewInt(1)
	return *res.Lsh(big.NewInt(int64(n)), uint(1))
}
