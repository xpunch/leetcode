package sorting

func countingSort(arr []int) []int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	counts := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}
	for i := 0; i < len(arr); {
		for j := 0; ; {
			if counts[j] > 0 {
				arr[i] = j
				counts[j]--
				i++
				break
			}
			j++
		}
	}
	return arr
}
