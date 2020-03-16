#
# @lc app=leetcode id=67 lang=python3
#
# [67] Add Binary
#

# @lc code=start


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        result = ''
        na, nb, isOverflow = len(a), len(b), False
        for i in range(max(na, nb)):
            bsum = 0
            if isOverflow:
                bsum += 1
            if i < na:
                if a[na-i-1] == '1':
                    bsum += 1
            if i < nb:
                if b[nb-i-1] == '1':
                    bsum += 1
            if bsum == 3:
                result = '1'+result
                isOverflow = True
            elif bsum == 2:
                result = '0'+result
                isOverflow = True
            elif bsum == 1:
                result = '1'+result
                isOverflow = False
            else:
                result = '0'+result
                isOverflow = False
        if isOverflow:
            result = '1'+result
        return result
# @lc code=end
