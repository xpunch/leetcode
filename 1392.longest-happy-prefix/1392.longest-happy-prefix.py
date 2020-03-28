#
# @lc app=leetcode id=1392 lang=python3
#
# [1392] Longest Happy Prefix
#

# @lc code=start


class Solution:
    def longestPrefix(self, s: str) -> str:
        n = len(s)
        if n <= 1:
            return ""
        for i in range(n-1, 0, -1):
            if s[:i] == s[n-i:]:
                return s[:i]
        return ""
# @lc code=end
