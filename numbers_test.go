// FILEPATH: /home/aho/git/go_katas/numbers_test.go

package go_kata

import (
	"testing"
)

func TestShiftLeft(t *testing.T) {
	testCases := []struct {
		n        int
		i        uint
		expected int
	}{
		{2, 1, 4},
		{3, 1, 6},
		{2, 2, 8},
		{5, 2, 20},
		{10, 3, 80},
		{0, 25, 0},
		{1024, 1, 2048},
		{-5, 1, -10},
	}

	for _, testCases := range testCases {
		result := ShiftLeft(testCases.n, testCases.i)
		if result != testCases.expected {
			t.Errorf("shiftLeft(%d, %d) = %d, expected %d", testCases.n, testCases.i, result, testCases.expected)
		}
	}
}

func TestShiftLeftBig(t *testing.T) {
	testCases := []struct {
		n        uint
		i        uint
		expected int
	}{
		{2, 1, 4},
		{3, 1, 6},
		{2, 2, 8},
		{5, 2, 20},
		{10, 3, 80},
		{0, 25, 0},
		{1024, 1, 2048},
		{1, 10, 1024},
	}

	for _, testCases := range testCases {
		result := ShiftLeftBig(testCases.n, testCases.i)
		if result.Int64() != int64(testCases.expected) {
			t.Errorf("shiftLeftBig(%d, %d) = %d, expected %d", testCases.n, testCases.i, result.Uint64(), testCases.expected)
		}
	}
}

func TestShiftLeftPow2(t *testing.T) {

	testCases := []struct {
		n        uint
		expected int
	}{
		{2, 4},
		{3, 6},
		{5, 10},
		{10, 20},
		{0, 0},
		{1024, 2048},
	}

	for _, testCases := range testCases {
		result := ShiftLeftMultBy2(testCases.n)
		if result.Int64() != int64(testCases.expected) {
			t.Errorf("shiftLeftBig(%d) = %d, expected %d", testCases.n, result.Uint64(), testCases.expected)
		}
	}
}
