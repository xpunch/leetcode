#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] Longest Substring Without Repeating Characters
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        max = 1
        start = 0
        cache = {s[0]:0}
        for i in range(1, len(s)):
            if cache.__contains__(s[i]):
                if cache[s[i]] >= start:
                    start = cache[s[i]] + 1
                    cache[s[i]] = i
                    continue
            if i - start + 1 > max:
                max = i - start + 1
            cache[s[i]] = i
        return max
# @lc code=end

