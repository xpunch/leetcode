# Edit Distance
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Hard (41.39%)|3080|45|

## Tags
string | dynamic-programming

## Companies
Unknown
```
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
```
## Example 1:
```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
```
## Example 2:
```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
```

## Submission
Python
```
√ Accepted
  √ 1146/1146 cases passed (204 ms)
  √ Your runtime beats 40.55 % of python3 submissions
  √ Your memory usage beats 15.38 % of python3 submissions (17.3 MB)
```