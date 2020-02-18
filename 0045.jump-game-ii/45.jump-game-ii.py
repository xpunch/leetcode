#
# @lc app=leetcode id=45 lang=python3
#
# [45] Jump Game II
#

# @lc code=start


class Solution:
    def jump(self, nums: List[int]) -> int:
        n, jumps = len(nums), 0
        if n <= 1:
            return 0
        if nums[0] >= n-1:
            return 1
        start, end, fastest = 0, 0, 0
        while fastest < n-1:
            for i in range(start, end+1):
                if(i+nums[i] > fastest):
                    fastest = i+nums[i]
            start = end+1
            end = fastest
            jumps += 1
        return jumps
# @lc code=end
