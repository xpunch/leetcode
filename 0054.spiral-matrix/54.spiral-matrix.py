#
# @lc app=leetcode id=54 lang=python3
#
# [54] Spiral Matrix
#
from typing import List

# @lc code=start


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        result = []
        m = len(matrix)
        if m == 0:
            return result
        n = len(matrix[0])
        row, column, i, j = 0, 0, 0, 0
        left, right, top, bottom = 0, n-1, 0, m-1
        direction = 0  # , 1 ↓, 2 ←, 3 ↑
        while len(result) < m*n:
            result.append(matrix[row][column])
            if direction == 0:  # →
                if column < right:
                    column += 1
                else:
                    direction += 1
                    row += 1
                    top += 1
            elif direction == 1:  # ↓
                if row < bottom:
                    row += 1
                else:
                    direction += 1
                    column -= 1
                    right -= 1
            elif direction == 2:  # ←
                if column > left:
                    column -= 1
                else:
                    direction += 1
                    row -= 1
                    bottom -= 1
            else:  # ↑
                if row > top:
                    row -= 1
                else:
                    direction = 0
                    column += 1
                    left += 1
        return result
# @lc code=end
