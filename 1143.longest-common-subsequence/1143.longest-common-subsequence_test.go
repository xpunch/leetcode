package xcode

import "testing"

func Test1(t *testing.T) {
	text1, text2 := "abcde", "ace"
	if output := longestCommonSubsequence(text1, text2); output != 3 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	text1, text2 := "abc", "abc"
	if output := longestCommonSubsequence(text1, text2); output != 3 {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	text1, text2 := "abc", "def"
	if output := longestCommonSubsequence(text1, text2); output != 0 {
		t.Fatal(output)
	}
}

func Test12(t *testing.T) {
	text1, text2 := "bl", "yby"
	if output := longestCommonSubsequence(text1, text2); output != 1 {
		t.Fatal(output)
	}
}

func Test15(t *testing.T) {
	text1, text2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	if output := longestCommonSubsequence(text1, text2); output != 210 {
		t.Fatal(output)
	}
}
