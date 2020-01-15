#
# @lc app=leetcode id=1074 lang=python3
#
# [1074] Number of Submatrices That Sum to Target
#

# @lc code=start


class Solution:

    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        x = len(matrix)
        y = len(matrix[0])
        sumCache = [[0 for j in range(y)] for i in range(x)]
        result = 0
        for i in range(x):
            for j in range(y):
                sumCache[i][j] = matrix[i][j]
                if i > 0:
                    sumCache[i][j] += sumCache[i-1][j]
                if j > 0:
                    sumCache[i][j] += sumCache[i][j-1]
                if i > 0 and j > 0:
                    sumCache[i][j] -= sumCache[i-1][j-1]
        for i in range(x):
            for k in range(i+1):
                rowCache = {}
                for j in range(y):
                    s = sumCache[i][j]
                    if k > 0:
                        s -= sumCache[k-1][j]
                    if s == target:
                        result += 1
                    if s-target in rowCache:
                        result += rowCache[s-target]
                    rowCache[s] = rowCache.setdefault(s, 0)+1
        return result
# @lc code=end
