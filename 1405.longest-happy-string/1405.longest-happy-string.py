#
# @lc app=leetcode id=1405 lang=python3
#
# [1405] Longest Happy String
#

# @lc code=start


class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        return self.build(a, b, c, '')

    def build(self, a: int, b: int, c: int, s: str) -> str:
        if a == 0 and b == 0 and c == 0:
            return s
        if s == '':
            if a >= b and a >= c:
                if self.canBeTwo(a, b, c):
                    s += 'aa'
                    return self.build(a-2, b, c, s)
                else:
                    s += 'a'
                    return self.build(a-1, b, c, s)
            elif b >= a and b >= c:
                if self.canBeTwo(b, a, c):
                    s += 'bb'
                    return self.build(a, b-2, c, s)
                else:
                    s += 'b'
                    return self.build(a, b-1, c, s)
            else:
                if self.canBeTwo(c, a, b):
                    s += 'cc'
                    return self.build(a, b, c-2, s)
                else:
                    s += 'c'
                    return self.build(a, b, c-1, s)
        elif s[-1] == 'a':
            if b == 0 and c == 0:
                return s
            if b >= c:
                if self.canBeTwo(b, a, c):
                    s += 'bb'
                    return self.build(a, b-2, c, s)
                else:
                    s += 'b'
                    return self.build(a, b-1, c, s)
            else:
                if self.canBeTwo(c, a, b):
                    s += 'cc'
                    return self.build(a, b, c-2, s)
                else:
                    s += 'c'
                    return self.build(a, b, c-1, s)
        elif s[-1] == 'b':
            if a == 0 and c == 0:
                return s
            if a >= c:
                if self.canBeTwo(a, b, c):
                    s += 'aa'
                    return self.build(a-2, b, c, s)
                else:
                    s += 'a'
                    return self.build(a-1, b, c, s)
            else:
                if self.canBeTwo(c, a, b):
                    s += 'cc'
                    return self.build(a, b, c-2, s)
                else:
                    s += 'c'
                    return self.build(a, b, c-1, s)
        elif s[-1] == 'c':
            if a == 0 and b == 0:
                return s
            if a >= b:
                if self.canBeTwo(a, b, c):
                    s += 'aa'
                    return self.build(a-2, b, c, s)
                else:
                    s += 'a'
                    return self.build(a-1, b, c, s)
            else:
                if self.canBeTwo(b, a, c):
                    s += 'bb'
                    return self.build(a, b-2, c, s)
                else:
                    s += 'b'
                    return self.build(a, b-1, c, s)

    def getMinSepartors(self, x: int) -> int:
        if x <= 2:
            return 0
        return (x-1)//2

    def canBeTwo(self, x: int, m: int, n: int) -> bool:
        return x >= 2 and x//2 > self.getMinSepartors(m)+self.getMinSepartors(n)
# @lc code=end
