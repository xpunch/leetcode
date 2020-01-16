# Brick Wall
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Medium (48.67%)|683|35|

## Tags
hash-table

## Companies
facebook

```
There is a brick wall in front of you. The wall is rectangular and has several rows of bricks. The bricks have the same height but different width. You want to draw a vertical line from the top to the bottom and cross the least bricks.

The brick wall is represented by a list of rows. Each row is a list of integers representing the width of each brick in this row from left to right.

If your line go through the edge of a brick, then the brick is not considered as crossed. You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.

You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.
```

## Example:
```
Input: [[1,2,2,1],
        [3,1,2],
        [1,3,2],
        [2,4],
        [3,1,2],
        [1,3,1,1]]

Output: 2
```
## Explanation:
![avatar](brick_wall.png)

## Note:
```
The width sum of bricks in different rows are the same and won't exceed INT_MAX.
The number of bricks in each row is in range [1,10,000]. The height of wall is in range [1,10,000]. Total number of bricks of the wall won't exceed 20,000.
```

## Python
```
√ Accepted
  √ 85/85 cases passed (180 ms)
  √ Your runtime beats 97.16 % of python3 submissions
  √ Your memory usage beats 50 % of python3 submissions (17.4 MB)
```