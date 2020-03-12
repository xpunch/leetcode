#
# @lc app=leetcode id=64 lang=python3
#
# [64] Minimum Path Sum
#
from typing import List
# @lc code=start


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if n == 0:
            return 0
        m = len(grid[0])
        dp = [[0 for j in range(m)] for i in range(n)]
        for i in range(n):
            for j in range(m):
                if i == 0 and j == 0:
                    dp[i][j] = grid[i][j]
                elif i == 0 and j > 0:
                    dp[i][j] = grid[i][j]+dp[i][j-1]
                elif i > 0 and j == 0:
                    dp[i][j] = grid[i][j]+dp[i-1][j]
                else:
                    dp[i][j] = grid[i][j]+min(dp[i][j-1], dp[i-1][j])
        return dp[n-1][m-1]
# @lc code=end
