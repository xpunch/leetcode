#
# @lc app=leetcode id=36 lang=python3
#
# [36] Valid Sudoku
#

# @lc code=start


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            rowCache = {}
            columnCache = {}
            cellCache = {}
            for j in range(9):
                num = board[i][j]
                if num != '.':
                    if num in rowCache:
                        return False
                    else:
                        rowCache[num] = 0
                num = board[j][i]
                if num != '.':
                    if num in columnCache:
                        return False
                    else:
                        columnCache[num] = 0
                row = j//3+(i//3)*3
                column = j % 3+(i % 3)*3
                num = board[row][column]
                if num != '.':
                    if num in cellCache:
                        return False
                    else:
                        cellCache[num] = 0
        return True
# @lc code=end
