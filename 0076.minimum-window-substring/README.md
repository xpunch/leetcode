# Minimum Window Substring
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Hard (33.61%)|3741|261|

## Tags
hash-table | two-pointers | string | sliding-window

## Companies
facebook | linkedin | snapchat | uber
```
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
```
## Example:
```
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```
## Note:

* If there is no such window in S that covers all characters in T, return the empty string "".
* If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## Notes:
```
Use array instead of map, cause characters are ascii.
```

## Submission
Python
```
√ Accepted
  √ 268/268 cases passed (84 ms)
  √ Your runtime beats 95.77 % of python3 submissions
  √ Your memory usage beats 5.55 % of python3 submissions (14.6 MB)
```