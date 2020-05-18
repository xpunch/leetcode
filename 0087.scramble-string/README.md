# Scramble String
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Hard (32.99%)|443|661|

## Tags
string | dynamic-programming

## Companies
Unknown
```
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
```
## Example 1:
```
Input: s1 = "great", s2 = "rgeat"
Output: true
```
## Example 2:
```
Input: s1 = "abcde", s2 = "caebd"
Output: false
```

## Submission
```
√ Accepted
  √ 283/283 cases passed (8 ms)
  √ Your runtime beats 24 % of golang submissions
  √ Your memory usage beats 100 % of golang submissions (3.8 MB)
```