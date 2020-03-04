#
# @lc app=leetcode id=49 lang=python3
#
# [49] Group Anagrams
#

# @lc code=start


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        cache, result = {}, []
        for i in range(len(strs)):
            k = "".join(sorted(list(strs[i])))
            if cache.__contains__(k):
                cache[k].append(strs[i])
            else:
                cache[k] = [strs[i]]
        for k in cache:
            result.append(cache[k])
        return result
# @lc code=end
