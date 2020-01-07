#
# @lc app=leetcode id=17 lang=python3
#
# [17] Letter Combinations of a Phone Number
#

# @lc code=start


class Solution:

    maps = {}
    
    def __init__(self):
        self.maps['2'] = ['a', 'b', 'c']
        self.maps['3'] = ['d', 'e', 'f']
        self.maps['4'] = ['g', 'h', 'i']
        self.maps['5'] = ['j', 'k', 'l']
        self.maps['6'] = ['m', 'n', 'o']
        self.maps['7'] = ['p', 'q', 'r', 's']
        self.maps['8'] = ['t', 'u', 'v']
        self.maps['9'] = ['w', 'x', 'y', 'z']

    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0:
            return []
        result = []
        self.findCombination(digits, 0, '', result)
        return result

    def findCombination(self, digits: str, index: int, letter:str, result: [str]):
        if index == len(digits):
            result.append(letter)
        else:
            tmp = self.maps[digits[index]]
            for i in range(len(tmp)):
                self.findCombination(digits, index+1, letter+tmp[i], result)
# @lc code=end
