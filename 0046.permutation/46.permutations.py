#
# @lc app=leetcode id=46 lang=python3
#
# [46] Permutations
#

# @lc code=start

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n, result = len(nums), []
        if n == 1:
            return [nums]
        index, indexs, rollback = n-2, [0 for i in range(n)], False
        while True:
            if rollback:
                if indexs[index] < n-index-1:
                    indexs[index] += 1
                    index = n-2
                    rollback = False
                else:
                    if index == 0:
                        break
                    indexs[index] = 0
                    index -= 1
            else:
                cache, num = nums[:], [0 for i in range(n)]
                for i in range(n):
                    num[i] = cache[indexs[i]]
                    del cache[indexs[i]]
                result.append(num)
                if indexs[index] < n-index-1:
                    indexs[index] += 1
                else:
                    if index == 0:
                        break
                    rollback = True
                    for i in range(index, n-1):
                        indexs[i] = 0
                    index -= 1
        return result
# @lc code=end
