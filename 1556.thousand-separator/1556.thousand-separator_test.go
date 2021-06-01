package xcode

import "testing"

func Test1(t *testing.T) {
	input := 987
	if output := thousandSeparator(input); output != "987" {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	input := 1234
	if output := thousandSeparator(input); output != "1.234" {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	input := 123456789
	if output := thousandSeparator(input); output != "123.456.789" {
		t.Fatal(output)
	}
}

func Test4(t *testing.T) {
	input := 0
	if output := thousandSeparator(input); output != "0" {
		t.Fatal(output)
	}
}
