# Rotate Image
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Medium (52.23%)|2408|201|

## Tags
array

## Companies
amazon | apple | microsoft
```
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).
```
## Note:
```
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
```
## Example 1:
```
Given input matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
```
## Example 2:
```
Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
```
## Example 3:
```
Given input matrix =
[
  [ 1, 2, 3, 4], 0,0 0,1 0,2 0,3
  [ 5, 6, 7, 8], 1,0 1,1 1,2 1,3
  [ 9,10,11,12], 2,0 2,1 2,2 2,3
  [13,14,15,16]  3,0 3,1 3,2 3,3
], 

rotate the input matrix in-place such that it becomes:
[
  [13, 9, 5, 1],
  [14,10, 6, 2],
  [15,11, 7, 3],
  [16,12, 8, 4]
]
```

## Python
```
√ Accepted
  √ 21/21 cases passed (32 ms)
  √ Your runtime beats 73.99 % of python3 submissions
  √ Your memory usage beats 100 % of python3 submissions (12.8 MB)
```