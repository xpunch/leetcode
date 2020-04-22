package main

import "testing"

func Test1(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	result := exist(board, word)
	if !result {
		t.Fatalf("%s %v", word, result)
	}
}

func Test2(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "SEE"
	result := exist(board, word)
	if !result {
		t.Fatalf("%s %v", word, result)
	}
}

func Test3(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "ABCB"
	result := exist(board, word)
	if result {
		t.Fatalf("%s %v", word, result)
	}
}
