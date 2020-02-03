#
# @lc app=leetcode id=39 lang=python3
#
# [39] Combination Sum
#

# @lc code=start


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        result, indexs, s, n = [], [-1], 0, len(candidates)
        while len(indexs) > 0:
            if indexs[-1] < n-1:
                indexs[-1] += 1
                s += candidates[indexs[-1]]
            else:
                s -= candidates[indexs[-1]]
                indexs.pop()
                if len(indexs) > 0 and indexs[-1] != n-1:
                    s -= candidates[indexs[-1]]
                continue
            if s < target:
                indexs.append(indexs[-1]-1)
            else:
                if s == target:
                    com = []
                    for i in range(len(indexs)):
                        com.append(candidates[indexs[i]])
                    result.append(com)
                s -= candidates[indexs[-1]]
                indexs.pop()
                if len(indexs) > 0 and indexs[-1] != n-1:
                    s -= candidates[indexs[-1]]
        return result
# @lc code=end
