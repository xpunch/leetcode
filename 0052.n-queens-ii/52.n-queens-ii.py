#
# @lc app=leetcode id=52 lang=python3
#
# [52] N-Queens II
#

# @lc code=start
from typing import List


class Solution:
    def totalNQueens(self, n: int) -> int:
        if n == 1:
            return 1
        result = 0
        for i in range(n):
            result += self.solve(n, [i], 0)
        return result

    def solve(self, n: int, queens: List[int], result: int) -> int:
        ap = self.getAvailablePositions(n, queens)
        if len(ap) == 0:
            return result
        row = len(queens)
        if row == n-1:
            return result+1
        else:
            for i in ap:
                tmp = queens[:]
                tmp.append(i)
                result += self.solve(n, tmp, 0)
            return result

    def getAvailablePositions(self, n: int, queens: List[int]) -> List[int]:
        row = len(queens)
        result = [i for i in range(n)]
        for i in range(row):
            queen = queens[i]
            gap = abs(row-i)
            if result.__contains__(queen):
                result.remove(queen)
            if queen - gap >= 0 and result.__contains__(queen-gap):
                result.remove(queen-gap)
            if queen+gap < n and result.__contains__(queen+gap):
                result.remove(queen+gap)
        return result
# @lc code=end
