package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test0(t *testing.T) {
	head := build([]int{1, 4, 3, 2, 5, 2})
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Print("->")
	result := partition(head, 3)
	for cur := result; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Println()
}

func Test1(t *testing.T) {
	head := build([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Print("->")
	result := partition(head, 4)
	for cur := result; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Println()
}

func Test2(t *testing.T) {
	head := build([]int{5, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Print("->")
	result := partition(head, 4)
	for cur := result; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Println()
}

func build(inputs []int) *ListNode {
	var root *ListNode
	if len(inputs) > 0 {
		root = &ListNode{Val: inputs[0], Next: build(inputs[1:])}
	}
	return root
}
