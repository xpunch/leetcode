#
# @lc app=leetcode id=1390 lang=python3
#
# [1390] Four Divisors
#

from typing import List
# @lc code=start


class Solution:
    def sumFourDivisors(self, nums: List[int]) -> int:
        result = 0
        for num in nums:
            divisors = self.getFourDivisors(num)
            if len(divisors)==4:
                result += (divisors[0]+divisors[1]+divisors[2]+divisors[3])
        return result

    def getFourDivisors(self, num: int) -> List[int]:
        if num < 6:
            return []
        result = [1]
        for i in range(2, num):
            if num //i < i:
                break
            if num % i == 0:
                p = num // i
                if p!=i:
                    if len(result)>1:
                        return []
                    result.append(i)
                    result.append(p)
                else:
                    if len(result)>2:
                        return []
                    result.append(i)
        result.append(num)
        return result
# @lc code=end

