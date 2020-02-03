#
# @lc app=leetcode id=38 lang=python3
#
# [38] Count and Say
#

# @lc code=start


class Solution:

    cache = {1: "1", 2: "11", 3: "21", 4: "1211", 5: "111221"}

    def countAndSay(self, n: int) -> str:
        if self.cache.__contains__(n):
            return self.cache[n]
        else:
            pre = self.countAndSay(n-1)
            result, count = "", 1
            for i in range(1, len(pre)):
                if pre[i] == pre[i-1]:
                    count += 1
                else:
                    result = result + str(count)+pre[i-1]
                    count = 1
            result = result + str(count)+pre[-1]
            self.cache[n] = result
            return result
            # @lc code=end
