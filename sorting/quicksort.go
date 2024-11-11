/*
Quick Sort is an divide-and-conquer sorting algorithm.

Steps:

1. Select pivot element.
2. Partition the array: all elements <= to the pivot are on its left, and all elements < on its right.
3. Apply recursively to the sub arrays until the entire array is sorted.

Complexity of O(n log n) and a worst-case complexity of O(n^2).
*/
package sorting

import "sync"

// Idiomatic implementation, uses slices. Turns out to have best performance!
func IdiomaticQuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Choose a pivot; here we take the middle element
	pivot := arr[len(arr)/2]

	// Partitions: left for elements less than pivot, right for greater than pivot
	// A slice is a lightweight descriptor that contains a pointer to the underlying array. passing a slice to a function, only this descriptor is copied.
	var left, right []int
	for _, element := range arr {
		if element <= pivot {
			right = append(right, element) //append is obviously very efficient in go.
		} else {
			left = append(left, element)
		}
	}
	return append(IdiomaticQuickSort(right), IdiomaticQuickSort(left)...)
}

// TODO, too much overhead in this implementation.
func ConcurrentQuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, right []int
	for _, element := range arr {
		if element <= pivot {
			right = append(right, element)
		} else {
			left = append(left, element)
		}
	}

	var leftSorted, rightSorted []int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		leftSorted = ConcurrentQuickSort(left)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rightSorted = ConcurrentQuickSort(right)
	}()

	wg.Wait() //blocks here util counter is back to 0

	// Concatenate leftSorted, pivot, and rightSorted
	return append(rightSorted, leftSorted...)
}
