#
# @lc app=leetcode id=59 lang=python3
#
# [59] Spiral Matrix II
#
from typing import List
# @lc code=start


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        matrix = [[0 for j in range(n)] for i in range(n)]
        direction, row, column = 0, 0, 0
        top, bottom, left, right = 0, n-1, 0, n-1
        for i in range(n*n):
            matrix[row][column] = i+1
            if direction == 0:
                if column < right:
                    column += 1
                else:
                    direction = 1
                    row += 1
                    top += 1
            elif direction == 1:
                if row < bottom:
                    row += 1
                else:
                    direction = 2
                    column -= 1
                    right -= 1
            elif direction == 2:
                if column > left:
                    column -= 1
                else:
                    direction = 3
                    row -= 1
                    bottom -= 1
            else:
                if row > top:
                    row -= 1
                else:
                    direction = 0
                    column += 1
                    left += 1
        return matrix
# @lc code=end
