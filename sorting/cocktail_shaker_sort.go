package sorting

/*
https://en.wikipedia.org/wiki/Cocktail_shaker_sort
Cocktail shaker sort, also known as bidirectional bubble sort, cocktail sort, shaker sort (which can also refer to a variant of selection sort), ripple sort, shuffle sort,[3] or shuttle sort,
is a variation of bubble sort that is both a stable sorting algorithm and a comparison sort.
The algorithm differs from a bubble sort in that it sorts in both directions on each pass through the list.
This sorting algorithm is only marginally more difficult to implement than a bubble sort, and solves the problem of turtles in bubble sorts.
It provides only marginal performance improvements, and does not improve asymptotic performance; like the bubble sort, it is not of practical interest (insertion sort is preferred for simple sorts), though it finds some use in education.
*/
func cocktailShakerSort(arr []int) []int {
	for {
		swapped := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
