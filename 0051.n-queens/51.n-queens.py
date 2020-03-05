#
# @lc app=leetcode id=51 lang=python3
#
# [51] N-Queens
#
from typing import List

# @lc code=start


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        if n == 1:
            return['Q']
        result = []
        for i in range(n):
            queens = [i]
            self.solve(n, 1, queens[:], result)
        return result

    def solve(self, n: int, row: int, queens: List[int], result: List[List[str]]):
        ap = self.getAvailablePositions(n, row, queens)
        if len(ap) == 0:
            return
        if row == n-1:
            queens.append(ap[0])
            board = []
            for i in range(n):
                line = ''
                for j in range(n):
                    if j != queens[i]:
                        line += '.'
                    else:
                        line += 'Q'
                board.append(line)
            result.append(board)
        else:
            for p in ap:
                tmp = queens[:]
                tmp.append(p)
                self.solve(n, row+1, tmp, result)

    def getAvailablePositions(self, n: int, row: int, queens: List[int]) -> List[int]:
        result = [i for i in range(n)]
        for i in range(row):
            q = queens[i]
            gap = abs(row-i)
            if result.__contains__(q):
                result.remove(q)
            if q-gap >= 0 and result.__contains__(q-gap):
                result.remove(q-gap)
            if q+gap < n and result.__contains__(q+gap):
                result.remove(q+gap)
        return result
# @lc code=end
