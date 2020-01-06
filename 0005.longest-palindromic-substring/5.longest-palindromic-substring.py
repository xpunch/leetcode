#
# @lc app=leetcode id=5 lang=python3
#
# [5] Longest Palindromic Substring
#

# @lc code=start


class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        dp = [[False for i in range(n)] for i in range(n)]
        l, r = 0, 0
        for j in range(1, n):
            for i in range(j-1, -1, -1):
                if (s[i] == s[j]) and (j-i <= 2 or dp[i+1][j-1]):
                    dp[i][j] = True
                    if j-i > r-l:
                        l = i
                        r = j
        return s[l:r+1]
# @lc code=end
