# Word Search
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Medium (33.94%)|3074|155|

## Tags
array | backtracking

## Companies
bloomberg | facebook | microsoft

```
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
```

## Example:
```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
```

## Constraints:
* board and word consists only of lowercase and uppercase English letters.
* 1 <= board.length <= 200
* 1 <= board[i].length <= 200
* 1 <= word.length <= 10^3

## Submission
Golang
```
√ Accepted
  √ 89/89 cases passed (4 ms)
  √ Your runtime beats 94.64 % of golang submissions
  √ Your memory usage beats 50 % of golang submissions (3.7 MB)
```