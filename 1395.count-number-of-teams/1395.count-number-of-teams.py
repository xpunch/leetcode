#
# @lc app=leetcode id=1395 lang=python3
#
# [1395] Count Number of Teams
#

# @lc code=start
class Solution:
    def numTeams(self, rating: List[int]) -> int:
        n = len(rating)
        if n < 3:
            return 0
        result = 0
        for i in range(n-2):
            s1 = rating[i]
            for j in range(i+1, n-1):
                s2 = rating[j]
                if s1 == s2:
                    continue
                elif s1 > s2:
                    for k in range(j+1, n):
                        s3 = rating[k]
                        if s2 > s3:
                            result += 1
                        else:
                            continue
                else:
                    for k in range(j+1, n):
                        s3 = rating[k]
                        if s2 < s3:
                            result += 1
                        else:
                            continue
        return result
# @lc code=end

