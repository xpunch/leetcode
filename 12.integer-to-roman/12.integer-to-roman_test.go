package main

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {
	roman := intToRoman(58)
	result := "LVIII"
	if roman != result {
		fmt.Printf("%s - %s", roman, result)
		t.Fail()
	}
}

func Test99(t *testing.T) {
	roman := intToRoman(1994)
	result := "MCMXCIV"
	if roman != result {
		fmt.Printf("%s - %s", roman, result)
		t.Fail()
	}
}
