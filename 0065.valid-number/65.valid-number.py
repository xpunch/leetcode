#
# @lc app=leetcode id=65 lang=python3
#
# [65] Valid Number
#

# @lc code=start


class Solution:
    def isNumber(self, s: str) -> bool:
        canEnd, canBeValid = False, False
        canBeSign, canBePoint, canHavePoint, canHaveE, haveE = True, True, True, False, False
        target = s.strip()
        for c in target:
            if c == '-' or c == '+':
                if canBeSign:
                    canHaveE = False
                    canBeSign = False
                    canBePoint = True
                    canEnd = False
                else:
                    return False
            elif c == '.':
                if canBePoint and canHavePoint:
                    canHavePoint = False
                    canBeSign = False
                    canBePoint = False
                    canEnd = True
                else:
                    return False
            elif c == 'e':
                if canHaveE and not haveE:
                    haveE = True
                    canHaveE = False
                    canBeSign = True
                    canBePoint = False
                    canHavePoint = False
                    canEnd = False
                else:
                    return False
            elif ord(c) >= 48 and ord(c) <= 57:
                canBeSign = False
                canEnd = True
                canBeValid = True
                canBePoint = True
                if not haveE:
                    canHaveE = True
                canHaveE = True
            else:
                return False
        return canEnd and canBeValid
    # @lc code=end
