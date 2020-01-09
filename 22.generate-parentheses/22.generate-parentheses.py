#
# @lc app=leetcode id=22 lang=python3
#
# [22] Generate Parentheses
#

# @lc code=start


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        if n == 1:
            return ['()']
        if n == 2:
            return ['(())', '()()']
        result = []
        self.generate(n, n, '', result)
        return result

    def generate(self, opened: int, closed: int, s: str, res: []):
        if opened == 0 and closed == 0:
            res.append(s)
        else:
            if opened > 0:
                self.generate(opened-1, closed, s+'(', res)
            if closed > opened:
                self.generate(opened, closed-1, s+')', res)
# @lc code=end
