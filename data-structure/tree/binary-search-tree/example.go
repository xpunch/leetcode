package main

import "fmt"

// TreeNode represents binary search tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Traversal will return nodes in-order traversal(LDR)
func (t *TreeNode) Traversal() []int {
	if t == nil {
		return nil
	}
	var result []int
	if t.Left != nil {
		result = append(result, t.Left.Traversal()...)
	}
	result = append(result, t.Val)
	if t.Right != nil {
		result = append(result, t.Right.Traversal()...)
	}
	return result
}

// Insert will insert node into tree
func (t *TreeNode) Insert(val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}
	if val < t.Val {
		t.Left = t.Left.Insert(val)
	} else if val > t.Val {
		t.Right = t.Right.Insert(val)
	}
	return t
}

// Delete wil delete node in tree
func (t *TreeNode) Delete(val int) *TreeNode {
	if t == nil {
		return nil
	}
	if val < t.Val {
		t.Left = t.Left.Delete(val)
	} else if val > t.Val {
		t.Right = t.Right.Delete(val)
	} else {
		if t.Left != nil && t.Right != nil {
			target := t.Left
			for target.Right != nil {
				target = target.Right
			}
			t = t.Delete(target.Val)
			t.Val = target.Val
		} else {
			if t.Left != nil {
				t = t.Left
			}
			t = t.Right
		}
	}
	return t
}

/*
    4
   / \
  2   5
 / \
1   3
Lets define root as D, left tree as L and right tree as R, there'll be six traversal ways: DLR, DRL, LDR, LRD, RDL, RLD.
If we limit it must through left to right, there will preorder(DLR), inorder(LDR), postorder(LRD) and breadth-first.
*/
func main() {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5}}
	fmt.Printf("%v\n", root.Traversal())

	root.Insert(0)
	fmt.Printf("%v\n", root.Traversal())

	root.Insert(6)
	fmt.Printf("%v\n", root.Traversal())

	root.Delete(5)
	fmt.Printf("%v\n", root.Traversal())

	root.Delete(2)
	fmt.Printf("%v\n", root.Traversal())
}
