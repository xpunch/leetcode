package main

import "fmt"

// TreeNode represents binary search tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreorderTraversal will return nodes in root-left-right orders(DLR)
func (t *TreeNode) PreorderTraversal() []int {
	if t == nil {
		return nil
	}
	var result []int
	result = append(result, t.Val)
	if t.Left != nil {
		result = append(result, t.Left.PreorderTraversal()...)
	}
	if t.Right != nil {
		result = append(result, t.Right.PreorderTraversal()...)
	}
	return result
}

// InorderTraversal will return nodes in left-root-right orders(LDR)
func (t *TreeNode) InorderTraversal() []int {
	if t == nil {
		return nil
	}
	var result []int
	if t.Left != nil {
		result = append(result, t.Left.InorderTraversal()...)
	}
	result = append(result, t.Val)
	if t.Right != nil {
		result = append(result, t.Right.InorderTraversal()...)
	}
	return result
}

// PostorderTraversal will return nodes in left-right-root roders(LRD)
func (t *TreeNode) PostorderTraversal() []int {
	if t == nil {
		return nil
	}
	var result []int
	if t.Left != nil {
		result = append(result, t.Left.PostorderTraversal()...)
	}
	if t.Right != nil {
		result = append(result, t.Right.PostorderTraversal()...)
	}
	result = append(result, t.Val)
	return result
}

// BreadthFirst will return nodes in breadth-first way recursively
func (t *TreeNode) BreadthFirst() []int {
	var result []int
	for i := 1; i <= t.Height(); i++ {
		result = append(result, t.nodesOfLevel(i)...)
	}
	return result
}

func (t *TreeNode) nodesOfLevel(l int) []int {
	if t == nil {
		return nil
	}
	if l == 1 {
		return []int{t.Val}
	}
	var result []int
	if t.Left != nil {
		result = append(result, t.Left.nodesOfLevel(l-1)...)
	}
	if t.Right != nil {
		result = append(result, t.Right.nodesOfLevel(l-1)...)
	}
	return result
}

// Height will return binary tree height
func (t *TreeNode) Height() int {
	if t == nil {
		return 0
	}
	var lh, rh int
	if t.Left != nil {
		lh = t.Left.Height()
	}
	if t.Right != nil {
		rh = t.Right.Height()
	}
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

/*
    1
   / \
  2   3
 / \
4   5
Lets define root as D, left tree as L and right tree as R, there'll be six traversal ways: DLR, DRL, LDR, LRD, RDL, RLD.
If we limit it must through left to right, there will preorder(DLR), inorder(LDR), postorder(LRD) and breadth-first.
*/
func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Printf("DLR: %v\n", root.PreorderTraversal())
	fmt.Printf("LDR: %v\n", root.InorderTraversal())
	fmt.Printf("LRD: %v\n", root.PostorderTraversal())
	fmt.Printf("BFS: %v\n", root.BreadthFirst())
}
