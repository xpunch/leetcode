#
# @lc app=leetcode id=58 lang=python3
#
# [58] Length of Last Word
#

# @lc code=start


class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        result, counting = 0, False
        for i in range(len(s)-1, -1, -1):
            if s[i] != ' ':
                counting = True
                result += 1
            elif counting:
                break
        return result
# @lc code=end
