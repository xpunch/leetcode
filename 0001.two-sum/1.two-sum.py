#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        if len(nums) < 2:
            return []
        cache = {}
        for i in range(0,len(nums)):
            last = target - nums[i]
            if cache.__contains__(last):
                return [i, cache[last]]
            cache[nums[i]]=i
        return []


# @lc code=end
