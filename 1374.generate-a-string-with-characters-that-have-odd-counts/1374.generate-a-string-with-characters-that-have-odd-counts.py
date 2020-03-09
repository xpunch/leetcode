#
# @lc app=leetcode id=1374 lang=python3
#
# [1374] Generate a String With Characters That Have Odd Counts
#

# @lc code=start


class Solution:
    def generateTheString(self, n: int) -> str:
        result, alphabets, cache = '', 'abcdefghijklmnopqrstuvwxyz', {}
        usedCharCount, usedCharIndex, charCount = len(alphabets), 0, 0
        if n % 2 != 0:
            usedCharCount -= 1
        while charCount < n:
            ch = alphabets[usedCharIndex]
            if cache.__contains__(ch):
                result = result + ch+ch
                charCount += 2
            else:
                cache[ch] = True
                result = result+ch
                charCount += 1
            usedCharIndex = (usedCharIndex+1) % usedCharCount
        return result

# @lc code=end
