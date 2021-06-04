package xcode

import "testing"

func Test1(t *testing.T) {
	arr, m := []int{3, 5, 1, 2, 4}, 1
	if output := findLatestStep(arr, m); output != 4 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	arr, m := []int{3, 1, 5, 4, 2}, 2
	if output := findLatestStep(arr, m); output != -1 {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	arr, m := []int{1}, 1
	if output := findLatestStep(arr, m); output != 1 {
		t.Fatal(output)
	}
}

func Test4(t *testing.T) {
	arr, m := []int{2, 1}, 2
	if output := findLatestStep(arr, m); output != 2 {
		t.Fatal(output)
	}
}

func Test24(t *testing.T) {
	arr, m := []int{3, 2, 5, 6, 10, 8, 9, 4, 1, 7}, 3
	if output := findLatestStep(arr, m); output != 9 {
		t.Fatal(output)
	}
}
