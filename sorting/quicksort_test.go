package sorting

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

const arraySize int = 100000

func TestIdiomaticQuickSort(t *testing.T) {

	testArray := getSamples(false)
	expected := getSamples(true)

	start := time.Now()
	var result = IdiomaticQuickSort(testArray)
	elapsed := time.Since(start)
	fmt.Printf("IdiomaticQuickSort: Time for %d elements  %s ", arraySize, elapsed)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("IdiomaticQuickSort failed. Got %v, expected %v", result, expected)
	}
}

func TestQuickSort(t *testing.T) {

	testArray := getSamples(false)
	expected := getSamples(true)

	start := time.Now()
	QuickSort(testArray)
	elapsed := time.Since(start)
	fmt.Printf("Quick Sort: Time for %d elements  %s ", arraySize, elapsed)
	if !reflect.DeepEqual(testArray, expected) {
		t.Errorf("Quick Sort failed. Got %v, expected %v", testArray, expected)
	}
}

func TestQuickSortTuned(t *testing.T) {

	testArray := getSamples(false)
	expected := getSamples(true)

	start := time.Now()
	QuickSort(testArray)
	elapsed := time.Since(start)
	fmt.Printf("Quick Sort: Time for %d elements  %s ", arraySize, elapsed)
	if !reflect.DeepEqual(testArray, expected) {
		t.Errorf("Quick Sort failed. Got %v, expected %v", testArray, expected)
	}
}

// samples helper
func getSamples(increasing bool) []int {
	result := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		if increasing {
			result[i] = i
		} else {
			result[i] = arraySize - 1 - i
		}
	}
	return result
}
