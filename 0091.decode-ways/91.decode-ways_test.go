package leetcode

import "testing"

func Test0(t *testing.T) {
	s := "12"
	result := numDecodings(s)
	if result != 2 {
		t.Fatalf("'%s':%d", s, result)
	}
}
func Test1(t *testing.T) {
	s := "226"
	result := numDecodings(s)
	if result != 3 {
		t.Fatalf("'%s':%d", s, result)
	}
}
func Test2(t *testing.T) {
	s := "0"
	result := numDecodings(s)
	if result != 0 {
		t.Fatalf("'%s':%d", s, result)
	}
}
func Test3(t *testing.T) {
	s := "1"
	result := numDecodings(s)
	if result != 1 {
		t.Fatalf("'%s':%d", s, result)
	}
}

// @lc code=end
