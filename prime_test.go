// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package go_kata

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	result := IsPrime(0)
	expected := false
	if result != expected {
		t.Errorf("IsPrime(0) = %t; want %t", result, expected)
	}
}
