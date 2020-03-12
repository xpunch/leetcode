#
# @lc app=leetcode id=63 lang=python3
#
# [63] Unique Paths II
#
from typing import List
# @lc code=start


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        n = len(obstacleGrid)
        if n == 0:
            return 0
        m = len(obstacleGrid[0])
        dp = [[0 for i in range(m)] for j in range(n)]
        for i in range(n):
            for j in range(m):
                b = obstacleGrid[i][j]
                if b == 0:
                    if i == 0 and j == 0:
                        dp[i][j] = 1
                    elif i == 0 and j != 0:
                        dp[i][j] = dp[i][j-1]
                    elif i != 0 and j == 0:
                        dp[i][j] = dp[i-1][j]
                    else:
                        dp[i][j] = dp[i-1][j]+dp[i][j-1]
                else:
                    dp[i][j] = 0
        return dp[n-1][m-1]
# @lc code=end
