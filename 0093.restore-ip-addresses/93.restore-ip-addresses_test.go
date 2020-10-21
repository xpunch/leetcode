package leetcode

import "testing"

func Test1(t *testing.T) {
	s := "25525511135"
	expected := []string{"255.255.11.135", "255.255.111.35"}
	result := restoreIpAddresses(s)
	if !equal(result, expected) {
		t.Fatalf("Output: %v, Expected: %v\n", result, expected)
	}
}
func Test2(t *testing.T) {
	s := "0000"
	expected := []string{"0.0.0.0"}
	result := restoreIpAddresses(s)
	if !equal(result, expected) {
		t.Fatalf("Output: %v, Expected: %v\n", result, expected)
	}
}

func Test3(t *testing.T) {
	s := "1111"
	expected := []string{"1.1.1.1"}
	result := restoreIpAddresses(s)
	if !equal(result, expected) {
		t.Fatalf("Output: %v, Expected: %v\n", result, expected)
	}
}

func Test4(t *testing.T) {
	s := "010010"
	expected := []string{"0.10.0.10", "0.100.1.0"}
	result := restoreIpAddresses(s)
	if !equal(result, expected) {
		t.Fatalf("Output: %v, Expected: %v\n", result, expected)
	}
}

func Test5(t *testing.T) {
	s := "101023"
	expected := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}
	result := restoreIpAddresses(s)
	if !equal(result, expected) {
		t.Fatalf("Output: %v, Expected: %v\n", result, expected)
	}
}

func equal(target, source []string) bool {
	if len(target) != len(source) {
		return false
	}
	for _, t := range target {
		found := false
		for _, s := range source {
			if s == t {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	for _, s := range source {
		found := false
		for _, t := range target {
			if s == t {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
