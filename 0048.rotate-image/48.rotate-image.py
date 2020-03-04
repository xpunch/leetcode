#
# @lc app=leetcode id=48 lang=python3
#
# [48] Rotate Image
#
from typing import List
# @lc code=start


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        if n == 1:
            return
        for i in range(n//2):
            w = n-i*2
            for j in range(w-1):
                x, y = i, i+j
                matrix[y][n-1-x], tmp = matrix[x][y], matrix[y][n-1-x]
                matrix[n-1-x][n-1-y], tmp = tmp, matrix[n-1-x][n-1-y]
                matrix[n-1-y][x], tmp = tmp, matrix[n-1-y][x]
                matrix[x][y] = tmp
# @lc code=end
