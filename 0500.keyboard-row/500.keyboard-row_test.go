package xcode

import "testing"

func Test1(t *testing.T) {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	if output := findWords(words); !equal(output, []string{"Alaska", "Dad"}) {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	words := []string{"amk"}
	if output := findWords(words); !equal(output, []string{}) {
		t.Fatal(output)
	}
}

func Test3(t *testing.T) {
	words := []string{"adsdf", "sfd"}
	if output := findWords(words); !equal(output, []string{"adsdf", "sfd"}) {
		t.Fatal(output)
	}
}
func equal(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}
