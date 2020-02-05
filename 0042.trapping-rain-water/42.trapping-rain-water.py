#
# @lc app=leetcode id=42 lang=python3
#
# [42] Trapping Rain Water
#

# @lc code=start


class Solution:
    def trap(self, height: List[int]) -> int:
        n, pos, result = len(height), 0, 0
        while pos < n-2:
            l, r = pos, 0
            if pos < n-1 and height[pos] > height[pos+1]:
                r = l+1
            else:
                for i in range(l, n-1):
                    if height[i+1] < height[i]:
                        l = i
                        r = i + 1
                        pos = i
                        break
                    elif height[i+1] > height[i]:
                        pos = i+1
                        break
                    else:
                        pos = i
            if r == 0:
                continue
            for i in range(r, n-1):
                pos = i
                if height[i+1] > height[i]:
                    r = i+1
                elif height[i+1] < height[i]:
                    break
            if height[l] > height[r] and r < n-1:
                mr = height[r]
                for i in range(r+1, n):
                    if height[i] > mr:
                        r = i
                        pos = i
                        mr = height[i]
                    if height[i] > height[l]:
                        r = i
                        pos = i
                        break
            h = min(height[l], height[r])
            for i in range(l, r+1):
                if height[i] < h:
                    result += h-height[i]
        return result

# @lc code=end
