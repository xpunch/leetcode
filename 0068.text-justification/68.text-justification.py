#
# @lc app=leetcode id=68 lang=python3
#
# [68] Text Justification
#
from typing import List
# @lc code=start


class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        result, n, wordsCount, cur = [], len(words), 0, 0
        for i in range(n):
            w = words[i]
            minLength = wordsCount+len(w)+i-cur
            if minLength <= maxWidth:
                if i == n-1:
                    result.append(self.leftJustify(words[cur:n], maxWidth))
                elif minLength+len(words[i+1])+1 > maxWidth:
                    result.append(self.fullyJustify(words[cur:i+1], maxWidth))
                    wordsCount = 0
                    cur = i+1
                    continue
            wordsCount += len(w)
        return result

    def fullyJustify(self, words: List[str], maxWidth: int) -> str:
        spaces, slots, line = maxWidth, len(words)-1, ''
        for w in words:
            spaces -= len(w)
        if slots == 0:
            line = words[0]
            for _ in range(spaces):
                line += ' '
            return line
        base, more = spaces // slots, spaces % slots
        for i in range(len(words)):
            line += words[i]
            if i != len(words)-1:
                spaces = base
                if i < more:
                    spaces += 1
                for _ in range(spaces):
                    line += ' '
        return line

    def leftJustify(self, words: List[str], maxWidth: int) -> str:
        line = ''
        for i in range(len(words)):
            line += words[i]
            if i != len(words)-1:
                line += ' '
            else:
                spaces = maxWidth - len(line)
                for _ in range(spaces):
                    line += ' '
        return line
# @lc code=end
