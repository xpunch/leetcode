#
# @lc app=leetcode id=74 lang=python3
#
# [74] Search a 2D Matrix
#
from typing import List
# @lc code=start


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m = len(matrix)
        if m == 0:
            return False
        n = len(matrix[0])
        if n == 0:
            return False
        if matrix[0][0] > target or matrix[m-1][n-1] < target:
            return False
        row, l, r = 0, 0, m-1
        if matrix[l][0] == target or matrix[r][0] == target:
            return True
        while l < r:
            if matrix[r][0] < target:
                row = r
                break
            if r - l == 1:
                row = l
                break
            mid = (l+r)//2
            if matrix[mid][0] == target:
                return True
            elif matrix[mid][0] > target:
                r = mid
            else:
                l = mid
        if n == 1:
            return False
        l, r = 1, n-1
        if matrix[row][l] == target or matrix[row][r] == target:
            return True
        while l < r:
            if r-l <= 1:
                return False
            mid = (l+r)//2
            if matrix[row][mid] == target:
                return True
            elif matrix[row][mid] > target:
                r = mid
            else:
                l = mid
    # @lc code=end
