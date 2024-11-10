package sorting

/*
Bubble Sort is a simple comparison-based sorting algorithm.
Steps:

 1. Start at the beginning. Repeatedly compare each pair of adjacent elements and swaps them if they are in the
    wrong order.

 2. Repeats this process for the remaining 'unsorted portion' (n-1) of the array until no swaps are needed.

In effect every iteration "bubbles" the largest unsorted element to its correct position at the end of the array after each full iteration.

Complexity for all cases is of O(n^2). Not efficient, not recommended for usage for large datasets.
*/
func BubbleSort(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			// if no swap occurred on the whole series run, everything is sorted
			break
		}
	}
}
