#
# @lc app=leetcode id=73 lang=python3
#
# [73] Set Matrix Zeroes
#
from typing import List
# @lc code=start


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        m = len(matrix)
        if m == 0:
            return None
        n = len(matrix[0])
        if n == 0:
            return None
        colZero, rowZero = False, False
        for i in range(m):
            if matrix[i][0] == 0:
                colZero = True
                break
        for j in range(n):
            if matrix[0][j] == 0:
                rowZero = True
                break
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    matrix[0][j] = 0
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0
        if colZero:
            for i in range(m):
                matrix[i][0] = 0
        if rowZero:
            for j in range(n):
                matrix[0][j] = 0
# @lc code=end
