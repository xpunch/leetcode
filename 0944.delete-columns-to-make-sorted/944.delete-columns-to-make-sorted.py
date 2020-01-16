#
# @lc app=leetcode id=944 lang=python3
#
# [944] Delete Columns to Make Sorted
#

# @lc code=start


class Solution:
    def minDeletionSize(self, A: List[str]) -> int:
        result = 0
        for i in range(len(A[0])):
            for j in range(1, len(A)):
                if A[j][i] < A[j-1][i]:
                    result += 1
                    break
        return result
# @lc code=end
