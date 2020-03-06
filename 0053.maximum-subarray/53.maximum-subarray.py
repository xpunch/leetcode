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
        return self.divideAndConquer(nums, 0, n-1)

    def divideAndConquer(self, nums: List[int], left: int, right: int) -> int:
        if left >= right:
            return nums[left]
        mid = left + (right - left)//2
        lmax = self.divideAndConquer(nums, left, mid-1)
        rmax = self.divideAndConquer(nums, mid+1, right)
        mmax = nums[mid]
        tmp = mmax
        for i in range(mid-1, left-1, -1):
            tmp += nums[i]
            if tmp > mmax:
                mmax = tmp
        tmp = mmax
        for i in range(mid+1, right+1):
            tmp += nums[i]
            if tmp > mmax:
                mmax = tmp
        return max(mmax, max(lmax, rmax))
        # @lc code=end
