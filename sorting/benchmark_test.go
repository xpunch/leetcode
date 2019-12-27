package sorting

import "testing"

const (
	ArraySize = 1000
)

func BenchmarkBubbleSort(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		bubbleSort(arr)
	}
}

func BenchmarkCocktailShakerSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		cocktailShakerSort(arr)
	}
}

func BenchmarkCombSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		combSort(arr)
	}
}

func BenchmarkGnomeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		gnomeSort(arr)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		selectionSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		mergeSort(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := RandomArray(ArraySize)
		b.StartTimer()
		quickSort(arr)
	}
}
