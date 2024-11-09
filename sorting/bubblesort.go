package sorting

// O(n2)
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
