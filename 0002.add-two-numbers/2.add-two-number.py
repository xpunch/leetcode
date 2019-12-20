#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] add tow numbers
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        r, n, n1, n2 = None, None, l1, l2
        overflow = False
        while n1 != None or n2 != None or overflow:
            val = 0
            if n1 != None:
                val += n1.val
                n1 = n1.next
            if n2 != None:
                val += n2.val
                n2 = n2.next
            if overflow:
                val += 1
                overflow = False
            if val >=10:
                val -= 10
                overflow = True
            t = ListNode(val)
            if r == None:
                n = r = t
            else:
                n.next = t
                n = t
        return r
        
# @lc code=end

