# Substring with Concatenation of All Words

|Category|Difficulty|Likes|Dislikes|
|-|-|-|-|
|algorithms|Hard (24.50%)|697|1081|

## Tags
hash-table | two-pointers | string

## Companies
Unknown

```
You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
```
 
## Example 1:
```
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
```
## Example 2:
```
Input:
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
Output: []
```

# Python
```
√ Accepted
  √ 173/173 cases passed (532 ms)
  √ Your runtime beats 39.22 % of python3 submissions
  √ Your memory usage beats 100 % of python3 submissions (12.9 MB)
```