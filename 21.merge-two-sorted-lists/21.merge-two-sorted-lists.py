#
# @lc app=leetcode id=21 lang=python3
#
# [21] Merge Two Sorted Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 == None:
            return l2
        if l2 == None:
            return l1
        if l1.val <= l2.val:
            l = self.mergeTwoLists(l1.next, l2)
            l1.next = l
            return l1
        else:
            l = self.mergeTwoLists(l1, l2.next)
            l2.next = l
            return l2
# @lc code=end

