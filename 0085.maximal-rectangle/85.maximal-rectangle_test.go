package main

import "testing"

func Test0(t *testing.T) {
	input := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	output := maximalRectangle(input)
	if output != 6 {
		t.Fatal(output)
	}
}
