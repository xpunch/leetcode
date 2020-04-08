#
# @lc app=leetcode id=76 lang=python3
#
# [76] Minimum Window Substring
#
# @lc code=start
from typing import Dict


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        n2 = len(t)
        if n2 == 0:
            return ""
        n1 = len(s)
        l, r, keys, last = 0, 0, [0 for i in range(128)], 0
        for c in t:
            keys[ord(c)] += 1
            last += 1
        start, end = -1, -1
        for i in range(n1):
            k = ord(s[i])
            keys[k] -= 1
            if keys[k] >= 0:
                last -= 1
            while last == 0:
                if (start == -1 and end == -1) or (i-l+1 < end-start+1):
                    start = l
                    end = i
                k = ord(s[l])
                keys[k] += 1
                if keys[k] > 0:
                    last += 1
                l += 1
        if start == -1 or end == -1:
            return ""
        return s[start:end+1]
# @lc code=end
