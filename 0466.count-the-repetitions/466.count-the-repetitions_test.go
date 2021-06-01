package xcode

import "testing"

func Test1(t *testing.T) {
	s1, n1, s2, n2 := "acb", 4, "ab", 2
	if output := getMaxRepetitions(s1, n1, s2, n2); output != 2 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	s1, n1, s2, n2 := "acb", 1, "acb", 1
	if output := getMaxRepetitions(s1, n1, s2, n2); output != 1 {
		t.Fatal(output)
	}
}

func Test4(t *testing.T) {
	s1, n1, s2, n2 := "aaa", 20, "aaaaa", 1
	if output := getMaxRepetitions(s1, n1, s2, n2); output != 12 {
		t.Fatal(output)
	}
}

func Test37(t *testing.T) {
	s1, n1, s2, n2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1000000, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1000000
	if output := getMaxRepetitions(s1, n1, s2, n2); output != 1 {
		t.Fatal(output)
	}
}
