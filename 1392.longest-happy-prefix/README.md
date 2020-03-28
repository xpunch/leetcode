# Longest Happy Prefix
|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Hard (37.70%)|115|10|

## Tags
Unknown

## Companies
Unknown

```
A string is called a happy prefix if is a non-empty prefix which is also a suffix (excluding itself).

Given a string s. Return the longest happy prefix of s .

Return an empty string if no such prefix exists.
```
 
## Example 1:
```
Input: s = "level"
Output: "l"
Explanation: s contains 4 prefix excluding itself ("l", "le", "lev", "leve"), and suffix ("l", "el", "vel", "evel"). The largest prefix which is also suffix is given by "l".
```
## Example 2:
```
Input: s = "ababab"
Output: "abab"
Explanation: "abab" is the largest prefix which is also suffix. They can overlap in the original string.
```
## Example 3:
```
Input: s = "leetcodeleet"
Output: "leet"
```
## Example 4:
```
Input: s = "a"
Output: ""
```

## Constraints:

* 1 <= s.length <= 10^5
* s contains only lowercase English letters.

## Submission
Python
```
√ Accepted
  √ 72/72 cases passed (1284 ms)
  √ Your runtime beats 56.46 % of python3 submissions
  √ Your memory usage beats 100 % of python3 submissions (15.1 MB)
```