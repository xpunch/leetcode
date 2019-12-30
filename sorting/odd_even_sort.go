package sorting

/*
https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort
In computing, an odd–even sort or odd–even transposition sort (also known as brick sort[self-published source]) is a relatively simple sorting algorithm, developed originally for use on parallel processors with local interconnections.
It is a comparison sort related to bubble sort, with which it shares many characteristics.
It functions by comparing all odd/even indexed pairs of adjacent elements in the list and, if a pair is in the wrong order (the first is larger than the second) the elements are switched.
The next step repeats this for even/odd indexed pairs (of adjacent elements). Then it alternates between odd/even and even/odd steps until the list is sorted.
*/
func oddEvenSort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}
