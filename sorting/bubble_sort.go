package sorting

/*
https://en.wikipedia.org/wiki/Bubble_sort
Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements and swaps them if they are in the wrong order.
The pass through the list is repeated until the list is sorted.
The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list.
Although the algorithm is simple, it is too slow and impractical for most problems even when compared to insertion sort.
Bubble sort can be practical if the input is in mostly sorted order with some out-of-order elements nearly in position.
*/
func bubbleSort(arr []int) []int {
	for i := 1; i <= len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
