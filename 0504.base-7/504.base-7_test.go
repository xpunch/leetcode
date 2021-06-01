package xcode

import "testing"

func Test1(t *testing.T) {
	input := 100
	if output := convertToBase7(input); output != "202" {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	input := -7
	if output := convertToBase7(input); output != "-10" {
		t.Fatal(output)
	}
}

func Test204(t *testing.T) {
	input := -999
	if output := convertToBase7(input); output != "-2625" {
		t.Fatal(output)
	}
}
