#
# @lc app=leetcode id=72 lang=python3
#
# [72] Edit Distance
#

# @lc code=start


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        n1, n2 = len(word1)+1, len(word2)+1
        dp = [[0 for j in range(n2)] for i in range(n1)]
        for i in range(1, n1):
            dp[i][0] = i
        for j in range(1, n2):
            dp[0][j] = j
        for i in range(1, n1):
            for j in range(1, n2):
                ops = 0
                if word1[i-1] != word2[j-1]:
                    ops = 1
                dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+ops)
        return dp[n1-1][n2-1]
# @lc code=end
