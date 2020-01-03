package sorting

/*
https://en.wikipedia.org/wiki/Shellsort
Shellsort, also known as Shell sort or Shell's method, is an in-place comparison sort.
It can be seen as either a generalization of sorting by exchange (bubble sort) or sorting by insertion (insertion sort).
The method starts by sorting pairs of elements far apart from each other, then progressively reducing the gap between elements to be compared.
Starting with far apart elements, it can move some out-of-place elements into position faster than a simple nearest neighbor exchange.
Donald Shell published the first version of this sort in 1959. The running time of Shellsort is heavily dependent on the gap sequence it uses.
For many practical variants, determining their time complexity remains an open problem.
*/
func shellSort(arr []int) []int {
	for g := len(arr) / 2; g > 0; g /= 2 {
		for i := g; i < len(arr); i++ {
			for j := i; j >= g && arr[j] < arr[j-g]; j -= g {
				arr[j], arr[j-g] = arr[j-g], arr[j]
			}
		}
	}
	return arr
}
