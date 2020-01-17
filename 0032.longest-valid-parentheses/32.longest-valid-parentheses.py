#
# @lc app=leetcode id=32 lang=python3
#
# [32] Longest Valid Parentheses
#

# @lc code=start


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        n, leftParts = len(s), []
        if n <= 1:
            return 0
        result, longth = 0, [0 for i in range(n)]
        if s[0] == '(':
            leftParts.append(0)
        for i in range(1, n):
            if s[i] == '(':
                leftParts.append(i)
                longth[i] = 0
            else:
                if len(leftParts) > 0:
                    j = leftParts.pop()
                    longth[i] = longth[j-1]+longth[i-1]+2
                else:
                    longth[i] = 0
            if longth[i] > result:
                result = longth[i]
        return result
# @lc code=end
