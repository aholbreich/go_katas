package sorting

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {

	//Approx the 1 mb array
	const arraySize int = 131072 * 1

	var arr [arraySize]int
	var expected [arraySize]int

	for i := 0; i < arraySize; i++ {
		arr[i] = arraySize - 1 - i
	}

	for i := 0; i < arraySize; i++ {
		expected[i] = i
	}

	start := time.Now()
	BubbleSort(arr[:])
	elapsed := time.Since(start)
	fmt.Printf("Bubble sort: Time for 100mb monotone int array %s\n", elapsed)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort failed. Got %v, expected %v", arr, expected)
	}
}
