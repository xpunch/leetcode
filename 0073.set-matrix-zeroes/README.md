# Set Matrix Zeroes
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Medium (42.13%)|1674|277|

## Tags
array

## Companies
amazon | microsoft
```
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
```
## Example 1:
```
Input: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
```
## Example 2:
```
Input: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
```
## Follow up:

* A straight forward solution using O(mn) space is probably a bad idea.
* A simple improvement uses O(m + n) space, but still not the best solution.
* Could you devise a constant space solution?

## Submission
Python
```
√ Accepted
  √ 159/159 cases passed (140 ms)
  √ Your runtime beats 50.8 % of python3 submissions
  √ Your memory usage beats 5.13 % of python3 submissions (14.1 MB)
```