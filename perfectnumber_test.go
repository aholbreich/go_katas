// FILEPATH: /home/aho/git/go_katas/perfectnumber_test.go

package go_kata

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFindPerfectNumbers(t *testing.T) {

	result := FindPerfectNumbers(4)
	expected := []int{6, 28, 496, 8128}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("findPerfectNumbers(%d) = %v, expected %v", 1, result, expected)
	}

}

func TestFindPerfectNumbersEfficient(t *testing.T) {

	result := FindPerfectNumbersEfficient(8)
	expected := []*big.Int{
		big.NewInt(6),
		big.NewInt(28),
		big.NewInt(496),
		big.NewInt(8128),
		big.NewInt(33550336),
		big.NewInt(8589869056),
		big.NewInt(137438691328),
		big.NewInt(2305843008139952128),
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FindPerfectNumbersEfficient(%d) = %v, expected %v", 1, result, expected)
	}
}
