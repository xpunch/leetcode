package xcode

import "testing"

func Test1(t *testing.T) {
	ops := []string{"5", "2", "C", "D", "+"}
	if output := calPoints(ops); output != 30 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	if output := calPoints(ops); output != 27 {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	ops := []string{"1"}
	if output := calPoints(ops); output != 1 {
		t.Fatal(output)
	}
}
