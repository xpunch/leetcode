package sorting

import "math/rand"

/*
https://en.wikipedia.org/wiki/Quicksort
Quicksort (sometimes called partition-exchange sort) is an efficient sorting algorithm, serving as a systematic method for placing the elements of a random access file or an array in order.
Developed by British computer scientist Tony Hoare in 1959 and published in 1961, it is still a commonly used algorithm for sorting.
When implemented well, it can be about two or three times faster than its main competitors, merge sort and heapsort.[contradictory]

Quicksort is a comparison sort, meaning that it can sort items of any type for which a "less-than" relation (formally, a total order) is defined.
Efficient implementations of Quicksort are not a stable sort, meaning that the relative order of equal sort items is not preserved.
Quicksort can operate in-place on an array, requiring small additional amounts of memory to perform the sorting.
It is very similar to selection sort, except that it does not always choose worst-case partition.

Mathematical analysis of quicksort shows that, on average, the algorithm takes O(n log n) comparisons to sort n items.
In the worst case, it makes O(n2) comparisons, though this behavior is rare.
*/
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := 0
	right := len(arr) - 1
	pivot := rand.Intn(len(arr))
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := 0; i < right; i++ {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}
