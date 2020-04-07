# Longest Happy String
| Category   | Difficulty      | Likes | Dislikes |
|------------|-----------------|-------|----------|
| algorithms | Medium (35.90%) | 36    | 29       |

## Tags
Unknown

## Companies
Unknown
```
A string is called happy if it does not have any of the strings 'aaa', 'bbb' or 'ccc' as a substring.

Given three integers a, b and c, return any string s, which satisfies following conditions:

s is happy and longest possible.
s contains at most a occurrences of the letter 'a', at most b occurrences of the letter 'b' and at most c occurrences of the letter 'c'.
s will only contain 'a', 'b' and 'c' letters.
If there is no such string s return the empty string "".
```

## Example 1:
```
Input: a = 1, b = 1, c = 7
Output: "ccaccbcc"
Explanation: "ccbccacc" would also be a correct answer.
```
## Example 2:
```
Input: a = 2, b = 2, c = 1
Output: "aabbc"
```
## Example 3:
```
Input: a = 7, b = 1, c = 0
Output: "aabaa"
Explanation: It's the only correct answer in this case.
```

## Constraints:

* 0 <= a, b, c <= 100
* a + b + c > 0

## Notes
```
'xxx' need to be sperated by 'y' or 'yy', so firstly found the sperated target, then found a valid sperator 'y' or 'yy'.
Biggest x used firstly, then the bigger one, if it's big enough then 'xx' otherwise 'x'.
```

## Submission
```
√ Accepted
  √ 33/33 cases passed (28 ms)
  √ Your runtime beats 100 % of python3 submissions
  √ Your memory usage beats 100 % of python3 submissions (13.9 MB)
```