#
# @lc app=leetcode id=66 lang=python3
#
# [66] Plus One
#
from typing import List
# @lc code=start


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        n, result = len(digits), digits[:]
        if result[-1] != 9:
            result[-1] += 1
        else:
            cur = n-1
            result[cur] += 1
            while result[cur] == 10:
                result[cur] = 0
                if cur == 0:
                    result.insert(0, 1)
                    break
                result[cur-1] += 1
                cur -= 1
        return result
# @lc code=end
