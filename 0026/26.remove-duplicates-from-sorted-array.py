#
# @lc app=leetcode id=26 lang=python3
#
# [26] Remove Duplicates from Sorted Array
#

# @lc code=start


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        n = len(nums)
        if n <= 1:
            return n
        i = 1
        while i < n:
            if nums[i] != nums[i-1]:
                i += 1
            else:
                del nums[i]
        return i
        # @lc code=end
