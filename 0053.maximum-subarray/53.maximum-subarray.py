#
# @lc app=leetcode id=53 lang=python3
#
# [53] Maximum Subarray
#

# @lc code=start


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 0:
            return 0
        if n == 1:
            return nums[0]
        result, tmp = nums[0], nums[0]
        for i in range(1, n):
            tmp = max(tmp+nums[i], nums[i])
            if tmp > result:
                result = tmp
        return result
        # @lc code=end
