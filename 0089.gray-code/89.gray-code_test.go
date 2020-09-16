package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test0(t *testing.T) {
	n := 0
	codes := grayCode(n)
	if len(codes) != 1<<n {
		t.Fail()
	}
	if codes[0] != 0 {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	n := 1
	codes := grayCode(n)
	if len(codes) != 1<<n {
		t.Fail()
	}
	if codes[0] != 0 {
		t.Fail()
	}
	if codes[1] != 1 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	n := 2
	codes := grayCode(n)
	if len(codes) != 1<<n {
		t.Fail()
	}
	format := "%0" + strconv.Itoa(n) + "b"
	for i := 0; i < (1<<n)-1; i++ {
		x, y := fmt.Sprintf(format, codes[i]), fmt.Sprintf(format, codes[i+1])
		var count int
		for j := 0; j < n; j++ {
			if x[j] != y[j] {
				count++
			}
		}
		if count > 1 {
			t.Fail()
		}
	}
}

func Test3(t *testing.T) {
	n := 3
	codes := grayCode(n)
	if len(codes) != 1<<n {
		t.Fail()
	}
	format := "%0" + strconv.Itoa(n) + "b"
	for i := 0; i < (1<<n)-1; i++ {
		x, y := fmt.Sprintf(format, codes[i]), fmt.Sprintf(format, codes[i+1])
		var count int
		for j := 0; j < n; j++ {
			if x[j] != y[j] {
				count++
			}
		}
		if count > 1 {
			t.Fail()
		}
	}
}

func Test4(t *testing.T) {
	n := 4
	codes := grayCode(n)
	if len(codes) != 1<<n {
		t.Fail()
	}
	format := "%0" + strconv.Itoa(n) + "b"
	for i := 0; i < (1<<n)-1; i++ {
		x, y := fmt.Sprintf(format, codes[i]), fmt.Sprintf(format, codes[i+1])
		var count int
		for j := 0; j < n; j++ {
			if x[j] != y[j] {
				count++
			}
		}
		if count > 1 {
			t.Fail()
		}
	}
}
