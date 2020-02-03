#
# @lc app=leetcode id=37 lang=python3
#
# [37] Sudoku Solver
#

# @lc code=start


class Solution:

    availables: map

    def __init__(self):
        self.availables = {}

    def solveSudoku(self, board: List[List[str]]) -> None:
        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    available = self.getAvilableValues(board, i, j)
                    if len(available) == 1:
                        board[i][j] = available[0]
                        self.refreshRelatedValues(board, i, j)
                    else:
                        self.availables[i*9+j] = available
        if len(self.availables) == 0:
            return None
        indexs = {}
        for k in self.availables:
            indexs[k] = 0
        keys, start = list(self.availables.keys()), 0
        while start < len(keys):
            i, j = keys[start]//9, keys[start] % 9
            available = self.getAvilableValues(board, i, j)
            if len(available) == 0:
                start -= 1
                indexs[keys[start]] += 1
                if board[i][j] != ".":
                    board[i][j] = "."
            else:
                if indexs[keys[start]] < len(available):
                    board[i][j] = available[indexs[keys[start]]]
                    start += 1
                else:
                    indexs[keys[start]] = 0
                    start -= 1
                    indexs[keys[start]] += 1
                    board[i][j] = "."

    def refreshRelatedValues(self, board: List[List[str]], i, j: int) -> None:
        for m in range(9):
            if m != i and self.availables.__contains__(m*9+j):
                available = self.availables[m*9+j]
                if board[i][j] in available:
                    available.remove(board[i][j])
                    if len(available) == 1:
                        board[m][j] = available[0]
                        del self.availables[m*9+j]
                        self.refreshRelatedValues(board, m, j)
        for n in range(9):
            if n != j and self.availables.__contains__(i*9+n):
                available = self.availables[i*9+n]
                if board[i][j] in available:
                    available.remove(board[i][j])
                    if len(available) == 1:
                        board[i][n] = available[0]
                        del self.availables[i*9+n]
                        self.refreshRelatedValues(board, i, n)
        for m in range(3):
            for n in range(3):
                k = (i//3)*3+m
                l = (j//3)*3+n
                if k != i and l != j and self.availables.__contains__(k*9+l):
                    available = self.availables[k*9+l]
                    if board[i][j] in available:
                        available.remove(board[i][j])
                        if len(available) == 1:
                            board[k][l] = available[0]
                            del self.availables[k*9+l]
                            self.refreshRelatedValues(board, k, l)

    def getAvilableValues(self, board: List[List[str]], i, j: int) -> List[str]:
        available = ['1', '2', '3', '4', '5', '6', '7', '8', '9']
        for n in range(9):
            if n != j and board[i][n] != "." and board[i][n] in available:
                available.remove(board[i][n])
        for m in range(9):
            if m != i and board[m][j] != "." and board[m][j] in available:
                available.remove(board[m][j])
        for m in range(3):
            for n in range(3):
                k = (i//3)*3+m
                l = (j//3)*3+n
                if k != i and l != j and board[k][l] != "." and board[k][l] in available:
                    available.remove(board[k][l])
        return available
# @lc code=end
