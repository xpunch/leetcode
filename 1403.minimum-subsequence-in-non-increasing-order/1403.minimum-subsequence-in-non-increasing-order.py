#
# @lc app=leetcode id=1403 lang=python3
#
# [1403] Minimum Subsequence in Non-Increasing Order
#

# @lc code=start


class Solution:
    def minSubsequence(self, nums: List[int]) -> List[int]:
        nums.sort(reverse=True)
        total = 0
        for n in nums:
            total += n
        sub = 0
        result = []
        for n in nums:
            if sub > total-sub:
                break
            else:
                sub += n
                result.append(n)
        return result
# @lc code=end

