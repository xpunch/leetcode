package xocde

import (
	"testing"
)

func Test1(t *testing.T) {
	input := []int{8, 1, 9}
	output := largestMultipleOfThree(input)
	if output != "981" {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	input := []int{8, 6, 7, 1, 0}
	output := largestMultipleOfThree(input)
	if output != "8760" {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	input := []int{1}
	output := largestMultipleOfThree(input)
	if output != "" {
		t.Fatal(output)
	}
}

func Test4(t *testing.T) {
	input := []int{0, 0, 0, 0, 0, 0}
	output := largestMultipleOfThree(input)
	if output != "0" {
		t.Fatal(output)
	}
}

func TestToString(t *testing.T) {
	if toString([]int{3, 0, 4, 5}, 3) != "540" {
		t.Fail()
	}
}
