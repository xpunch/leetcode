#
# @lc app=leetcode id=28 lang=python3
#
# [28] Implement strStr()
#

# @lc code=start


class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        l1 = len(needle)
        if l1 == 0:
            return 0
        l2 = len(haystack)
        if l2 < l1:
            return -1
        elif l2 == l1:
            if haystack == needle:
                return 0
            else:
                return -1
        for i in range(l2-l1+1):
            if haystack[i] == needle[0]:
                if l2 - i < l1:
                    return -1
                if haystack[i:i+l1] == needle:
                    return i
                if l2 - i == l1:
                    return -1
        return -1

# @lc code=end
