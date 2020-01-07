#
# @lc app=leetcode id=18 lang=python3
#
# [18] 4Sum
#

# @lc code=start


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        sum_cache = {}
        n = len(nums)
        nums.sort()
        for i in range(n):
            for j in range(i+1, n):
                s = nums[i]+nums[j]
                if sum_cache.__contains__(s):
                    sum_cache[s].append([i, j])
                else:
                    sum_cache[s] = [[i, j]]
        result = []
        for i in range(n):
            for j in range(i+1, n):
                last = target - nums[i] - nums[j]
                if not sum_cache.__contains__(last):
                    continue
                pairs = sum_cache[last]
                for p in pairs:
                    if p[0] == i or p[1] == i or p[0] == j or p[1] == j:
                        continue
                    quadruplet = [nums[i], nums[j], nums[p[0]], nums[p[1]]]
                    quadruplet.sort()
                    if not self.contains(quadruplet, result):
                        result.append(quadruplet)
        return result
    
    def contains(self, target: [int], array: [[int]]) -> bool:
        for arr in array:
            if target==arr:
                return True
        return False
# @lc code=end
