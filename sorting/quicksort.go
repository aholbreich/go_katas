/*
Quick Sort is an divide-and-conquer sorting algorithm.

Steps:

1. Select pivot element.
2. Partition the array: all elements <= to the pivot are on its left, and all elements < on its right.
3. Apply recursively to the sub arrays until the entire array is sorted.

Complexity of O(n log n) and a worst-case complexity of O(n^2).
*/
package sorting

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
		} else if element > pivot {
			left = append(left, element)
		}
	}
	return append(IdiomaticQuickSort(right), IdiomaticQuickSort(left)...)
}

// QuickSort sorts an array of integers using the Quick Sort algorithm. This implementation sorts on same array for efficiency.
// However this is slower than the Idiomatic version! Seem to be Golang specifics.
func QuickSort(arr []int) {

	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

// Recursive part
func quickSort(arr []int, lowIndex, highIndex int) {
	if lowIndex < highIndex {

		pivotIndex := partition(arr, lowIndex, highIndex)
		//recursion for both sides of pivot.
		quickSort(arr, lowIndex, pivotIndex-1)
		quickSort(arr, pivotIndex+1, highIndex)
	}
}

func partition(arr []int, lowIndex, highIndex int) int {
	// Choose the pivot. Here the last element as the pivot.
	pivot := arr[highIndex]
	i := lowIndex - 1 // Index of the smaller element

	for j := lowIndex; j < highIndex; j++ {
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place the pivot element in its correct position
	arr[i+1], arr[highIndex] = arr[highIndex], arr[i+1]
	return i + 1 // Return the pivot index
}

// Here i'm trying to Tune the one above, but it's only slightly better in results.
func QuickSortTuned(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivotIndex := partitionTuned(arr)

	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

func partitionTuned(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Place the pivot in its correct position
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i // Return the pivot index
}
