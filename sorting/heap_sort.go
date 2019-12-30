package sorting

/*
https://en.wikipedia.org/wiki/Heapsort
In computer science, heapsort is a comparison-based sorting algorithm.
Heapsort can be thought of as an improved selection sort: like that algorithm, it divides its input into a sorted and an unsorted region, and it iteratively shrinks the unsorted region by extracting the largest element and moving that to the sorted region.
The improvement consists of the use of a heap data structure rather than a linear-time search to find the maximum.

Although somewhat slower in practice on most machines than a well-implemented quicksort, it has the advantage of a more favorable worst-case O(n log n) runtime.
Heapsort is an in-place algorithm, but it is not a stable sort.

Heapsort was invented by J. W. J. Williams in 1964. This was also the birth of the heap, presented already by Williams as a useful data structure in its own right.
In the same year, R. W. Floyd published an improved version that could sort an array in-place, continuing his earlier research into the treesort algorithm.
*/
func heapSort(arr []int) []int {
	// 1. min heapify: adjust root node, make parent node smaller than children nodes
	// 2. build min heap
	// 3. remove root node, min heapify
	for i := 0; i < len(arr); i++ {
		minHeapify(arr[i:])
	}
	return arr
}

func minHeapify(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		if 2*i+2 < len(arr) && arr[2*i+2] < arr[2*i+1] {
			arr[2*i+1], arr[2*i+2] = arr[2*i+2], arr[2*i+1]
		}
		if arr[i] > arr[2*i+1] {
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		}
	}
}
