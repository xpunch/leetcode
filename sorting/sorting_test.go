package sorting

import (
	"math/rand"
	"testing"
	"time"
)

func TestBuddleSort(t *testing.T) {
	input := RandomArray(10)
	output := bubbleSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestSelectingSort(t *testing.T) {
	input := RandomArray(10)
	output := selectionSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestMergeSort(t *testing.T) {
	input := RandomArray(10)
	output := mergeSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestCocktailShakerSort(t *testing.T) {
	input := RandomArray(10)
	output := cocktailShakerSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestCombSort(t *testing.T) {
	input := RandomArray(100)
	output := combSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestGnomeSort(t *testing.T) {
	input := RandomArray(10)
	output := gnomeSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestQuickSort(t *testing.T) {
	input := RandomArray(10)
	output := quickSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestOddEvenSort(t *testing.T) {
	input := RandomArray(1000)
	output := oddEvenSort(input)
	if len(input) != len(output) {
		t.FailNow()
	}
	for i := 1; i < len(output); i++ {
		if output[i] < output[i-1] {
			t.FailNow()
		}
	}
}

func TestHeapSort(t *testing.T) {
	for c := 0; c < 1000; c++ {
		for i := 0; i < 10; i++ {
			input := RandomArray(10)
			output := heapSort(input)
			if len(input) != len(output) {
				t.FailNow()
			}
			for i := 1; i < len(output); i++ {
				if output[i] < output[i-1] {
					t.FailNow()
				}
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for c := 0; c < 1000; c++ {
		input := RandomArray(1000)
		output := insertionSort(input)
		if len(input) != len(output) {
			t.FailNow()
		}
		for i := 1; i < len(output); i++ {
			if output[i] < output[i-1] {
				t.FailNow()
			}
		}
	}
}

func TestShellSort(t *testing.T) {
	for c := 0; c < 1000; c++ {
		input := RandomArray(1000)
		output := shellSort(input)
		if len(input) != len(output) {
			t.FailNow()
		}
		for i := 1; i < len(output); i++ {
			if output[i] < output[i-1] {
				t.FailNow()
			}
		}
	}
}

func TestCountingSort(t *testing.T) {
	for c := 0; c < 1000; c++ {
		input := RandomArray(1000)
		output := countingSort(input)
		if len(input) != len(output) {
			t.FailNow()
		}
		for i := 1; i < len(output); i++ {
			if output[i] < output[i-1] {
				t.FailNow()
			}
		}
	}
}

func RandomArray(n int) []int {
	// needed a seed input else it will generate the same number
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
