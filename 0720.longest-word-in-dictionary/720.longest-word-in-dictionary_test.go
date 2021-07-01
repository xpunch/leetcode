package xcode

import "testing"

func Test1(t *testing.T) {
	words := []string{"w", "wo", "wor", "worl", "world"}
	if output := longestWord(words); output != "world" {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	if output := longestWord(words); output != "apple" {
		t.Fatal(output)
	}
}

func Test44(t *testing.T) {
	words := []string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"}
	if output := longestWord(words); output != "yodn" {
		t.Fatal(output)
	}
}
