#
# @lc app=leetcode id=1297 lang=python3
#
# [1297] Maximum Number of Occurrences of a Substring
#

# @lc code=start


class Solution:

    def maxFreq(self, s: str, maxLetters: int, minSize: int, maxSize: int) -> int:
        n = len(s)
        result = 0
        substrs = {}
        for i in range(n-minSize+1):
            cache = {}
            for j in range(i, min(i+maxSize, n)):
                letter = s[j]
                if cache.__contains__(letter):
                    cache[letter] += 1
                else:
                    cache[letter] = 1
                if len(cache) <= maxLetters and j-i+1 >= minSize:
                    substr = s[i:j+1]
                    if substrs.__contains__(substr):
                        substrs[substr] += 1
                    else:
                        substrs[substr] = 1
                    if substrs[substr] > result:
                        result = substrs[substr]
        return result
# @lc code=end
