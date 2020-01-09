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
        if n==2:
            return ['(())', '()()']
        result = []
        self.generate(n, 1, False, '(', result)
        return result

    def generate(self, n: int, open_count: int, rollbak: bool, s: str, res: []):
        if rollbak:
            if s == ')':
                return
            if s[-1] == '(':
                s = s[:-1]+')'
                self.generate(n, open_count-1, False, s, res)
            else:
                self.generate(n, open_count, True, s[:-1], res)
            return
        if len(s)==2*n-2:
            p = '('+s+')'
            if self.isValid(p):
                res.append(p)
            if s[-1] == '(':
                self.generate(n, open_count-1, True, s[:-1], res)
            else:
                self.generate(n, open_count, True, s[:-1], res)
        elif open_count < n-1:
            self.generate(n, open_count+1, False, s+'(', res)
        else:
            self.generate(n, open_count, False, s+')', res)
        pass

    def isValid(self, s: str) -> bool:
        cache = []
        for c in s:
            if len(cache) == 0:
                cache.append(c)
            else:
                last = cache[-1]
                if (last == '(' and c == ')') or (last == '[' and c == ']') or (last == '{' and c == '}'):
                    cache.pop()
                else:
                    cache.append(c)
        return len(cache) == 0
# @lc code=end
