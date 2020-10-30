# Unique Binary Search Trees II

| Category   | Difficulty      | Likes | Dislikes |
|------------|-----------------|-------|----------|
| algorithms | Medium (38.92%) | 2548  | 173      |

## Tags
dynamic-programming | tree

## Companies
Unknown
```
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
```
## Example:
```
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
```
## Explanation:
```
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

## Constraints:
```
0 <= n <= 8
```

## Submission
Golang Recursive
```
√ Accepted
  √ 9/9 cases passed (4 ms)
  √ Your runtime beats 94.64 % of golang submissions
  √ Your memory usage beats 11.61 % of golang submissions (5.2 MB)
```
Golang dynamic-programing
```
√ Accepted
  √ 9/9 cases passed (4 ms)
  √ Your runtime beats 94.64 % of golang submissions
  √ Your memory usage beats 11.61 % of golang submissions (4.8 MB)
```