package sorting

/*
In computer science, merge sort (also commonly spelled mergesort) is an efficient, general-purpose, comparison-based sorting algorithm.
Most implementations produce a stable sort, which means that the order of equal elements is the same in the input and output.
Merge sort is a divide and conquer algorithm that was invented by John von Neumann in 1945.
A detailed description and analysis of bottom-up mergesort appeared in a report by Goldstine and von Neumann as early as 1948.
*/
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]int, 0, len(arr))
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i >= len(left) {
			return append(result, right[j:]...)
		}
		if j >= len(right) {
			return append(result, left[i:]...)
		}
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return result
}
