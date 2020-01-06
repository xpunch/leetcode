#
# @lc app=leetcode id=16 lang=python3
#
# [16] 3Sum Closest
#

# @lc code=start


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        n = len(nums)
        nums.sort()
        closest = nums[0]+nums[1]+nums[2]
        for i in range(n-2):
            l, r = i + 1, n-1
            while l < r:
                s = nums[i]+nums[l]+nums[r]
                if abs(s-target) < abs(closest-target):
                    closest = s
                if s > target:
                    r -= 1
                else:
                    l += 1
        return closest
# @lc code=end
