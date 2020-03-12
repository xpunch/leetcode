# Unique Paths II
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Medium (33.94%)|1325|213|

## Tags
array | dynamic-programming

## Companies
bloomberg
```
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?
```
![img](robot_maze.png)
```
An obstacle and empty space is marked as 1 and 0 respectively in the grid.
```
**Note**: m and n will be at most 100.

## Example 1:
```
Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
```
## Explanation:
```
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
```

## Python
```
√ Accepted
  √ 43/43 cases passed (40 ms)
  √ Your runtime beats 91.07 % of python3 submissions
  √ Your memory usage beats 97.78 % of python3 submissions (12.9 MB)
```