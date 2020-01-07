#
# @lc app=leetcode id=17 lang=python3
#
# [17] Letter Combinations of a Phone Number
#

# @lc code=start


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0:
            return []
        maps = {}
        maps['2'] = ['a', 'b', 'c']
        maps['3'] = ['d', 'e', 'f']
        maps['4'] = ['g', 'h', 'i']
        maps['5'] = ['j', 'k', 'l']
        maps['6'] = ['m', 'n', 'o']
        maps['7'] = ['p', 'q', 'r', 's']
        maps['8'] = ['t', 'u', 'v']
        maps['9'] = ['w', 'x', 'y', 'z']
        n = len(digits)
        digit_size = [0 for i in range(n)]
        total = 1
        for i in range(n):
            digit_size[i] = len(maps[digits[i]])
            total *= digit_size[i]
        i = 0
        result = ['' for i in range(total)]
        for i in range(total):
            base = 1
            for j in range(1, n):
                base *= digit_size[j]
            sub = i
            tmp = [0 for i in range(n)]
            for j in range(n):
                tmp[j] = int(sub // base)
                if j != n-1:
                    sub %= base
                    base /= digit_size[j+1]
            letter = ''
            for j in range(n):
                letter += maps[digits[j]][tmp[j]]
            result[i] = letter
        return result
# @lc code=end
