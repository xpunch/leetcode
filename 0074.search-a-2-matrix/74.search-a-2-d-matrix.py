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
        row, l, r = 0, 0, m*n-1
        while l <= r:
            mid = l + ((r-l) >> 1)
            num = matrix[mid//n][mid % n]
            if num == target:
                return True
            elif num > target:
                r = mid-1
            else:
                l = mid+1
        return False
    # @lc code=end
