#
# @lc app=leetcode id=1389 lang=python3
#
# [1389] Create Target Array in the Given Order
#

# @lc code=start
class Solution:
    def createTargetArray(self, nums: List[int], index: List[int]) -> List[int]:
        n, result = len(nums), []
        for i in range(n):
            result.insert(index[i], nums[i])
        return result
# @lc code=end

