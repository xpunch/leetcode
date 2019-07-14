#Regular Expression Matching
|Category	|Difficulty	|Likes	|Dislikes|
-|-|-|-
|algorithms	|Hard (25.24%)	|2619	|512|
## Tags
string | dynamic-programming | backtracking

## Companies
airbnb | facebook | google | twitter | uber

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:
```
s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
```
Example 1:
```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```
Example 2:
```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```
Example 3:
```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```
Example 4:
```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
```
Example 5:
```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```
## Result
```
√ Accepted
  √ 447/447 cases passed (0 ms)
  √ Your runtime beats 100 % of golang submissions
  √ Your memory usage beats 57.2 % of golang submissions (2.2 MB)
```