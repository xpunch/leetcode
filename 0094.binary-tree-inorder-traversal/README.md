# [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/)

| Category   | Difficulty      | Likes | Dislikes |
|------------|-----------------|-------|----------|
| algorithms | Medium (61.32%) | 3760  | 161      |

## Tags
hash-table | stack | tree

## Companies
microsoft

```
Given the root of a binary tree, return the inorder traversal of its nodes' values.
```

## Example 1:
![inorder_1](inorder_1.jpg)
```
Input: root = [1,null,2,3]
Output: [1,3,2]
```
## Example 2:
```
Input: root = []
Output: []
```
## Example 3:
```
Input: root = [1]
Output: [1]
```
## Example 4:
![inorder_4](inorder_4.jpg)
```
Input: root = [1,2]
Output: [2,1]
```
## Example 5:
![inorder_5](inorder_5.jpg)
```
Input: root = [1,null,2]
Output: [1,2]
```

## Constraints:
```
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
```

## Follow up:
```
Recursive solution is trivial, could you do it iteratively?
```

## Submission
Go recursive
```
√ Accepted
  √ 68/68 cases passed (0 ms)
  √ Your runtime beats 100 % of golang submissions
  √ Your memory usage beats 100 % of golang submissions (2.1 MB)
```
Go iterative
```
√ Accepted
  √ 68/68 cases passed (0 ms)
  √ Your runtime beats 100 % of golang submissions
  √ Your memory usage beats 100 % of golang submissions (2 MB)
```