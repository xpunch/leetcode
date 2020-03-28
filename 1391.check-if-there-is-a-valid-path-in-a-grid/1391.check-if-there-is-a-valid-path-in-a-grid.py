#
# @lc app=leetcode id=1391 lang=python3
#
# [1391] Check if There is a Valid Path in a Grid
#
from typing import List
# @lc code=start


class Solution:
    def hasValidPath(self, grid: List[List[int]]) -> bool:
        m, n = len(grid), len(grid[0])
        # direction  0→ 1← 2↑ 3↓
        return self.isValid(m, n, grid, 0, 0, -1)

    def isValid(self, m: int, n: int, grid: List[List[int]], row: int, col: int, direction: int) -> bool:
        if row < 0 or row >= m or col < 0 or col >= n:
            return False
        if row == 0 and col == 0 and direction != -1:
            return False
        street = grid[row][col]
        print(row, col, street)
        if row == m-1 and col == n-1:
            if street == 1:
                if direction == -1 or direction == 0 or direction == 1:
                    return True
                else:
                    return False
            elif street == 2:
                if direction == -1 or direction == 3 or direction == 2:
                    return True
                else:
                    return False
            elif street == 3:
                if direction == -1 or direction == 0 or direction == 2:
                    return True
                else:
                    return False
            elif street == 4:
                if direction == -1 or direction == 1 or direction == 2:
                    return True
                else:
                    return False
            elif street == 5:
                if direction == 0 or direction == 3:
                    return True
                else:
                    return False
            else:
                if direction == -1 or direction == 3 or direction == 1:
                    return True
                else:
                    return False
        if street == 1:
            if direction == -1 or direction == 0:
                return self.isValid(m, n, grid, row, col+1, 0)
            elif direction == 1:
                return self.isValid(m, n, grid, row, col-1, 1)
            else:
                return False
        elif street == 2:
            if direction == -1 or direction == 3:
                return self.isValid(m, n, grid, row+1, col, 3)
            elif direction == 2:
                return self.isValid(m, n, grid, row-1, col, 2)
            else:
                return False
        elif street == 3:
            if direction == -1 or direction == 0:
                return self.isValid(m, n, grid, row+1, col, 3)
            elif direction == 2:
                return self.isValid(m, n, grid, row, col-1, 1)
            else:
                return False
        elif street == 4:
            if direction == -1:
                if self.isValid(m, n, grid, row, col+1, 0):
                    return True
                return self.isValid(m, n, grid, row+1, col, 3)
            elif direction == 1:
                return self.isValid(m, n, grid, row+1, col, 3)
            elif direction == 2:
                return self.isValid(m, n, grid, row, col+1, 0)
            else:
                return False
        elif street == 5:
            if direction == 0:
                return self.isValid(m, n, grid, row-1, col, 2)
            elif direction == 3:
                return self.isValid(m, n, grid, row, col-1, 1)
            else:
                return False
        else:
            if direction == -1 or direction == 3:
                return self.isValid(m, n, grid, row, col+1, 0)
            elif direction == 1:
                return self.isValid(m, n, grid, row-1, col, 2)
            else:
                return False
# @lc code=end
