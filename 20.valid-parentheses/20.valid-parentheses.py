#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#

# @lc code=start


class Solution:
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
