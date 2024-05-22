// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package go_kata

import (
	"testing"
)

func runIsPrimeTests(t *testing.T, testFunc func(int) bool) {
	tests := []struct {
		input    int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{7, true},
		{8, false},
		{71, true},
		{941, true},
		{1000, false},
		{7919, true},
	}
	for _, test := range tests {
		result := testFunc(test.input)
		if result != test.expected {
			t.Errorf("IsPrime(%d) = %t; want %t", test.input, result, test.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	runIsPrimeTests(t, IsPrime)
}

func TestIsPrimeUsingMathLib(t *testing.T) {
	runIsPrimeTests(t, IsPrimeWithMathLib)
}
