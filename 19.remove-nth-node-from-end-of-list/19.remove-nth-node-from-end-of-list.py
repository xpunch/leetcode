#
# @lc app=leetcode id=19 lang=python3
#
# [19] Remove Nth Node From End of List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        tail, pre_node = head, None
        i = 0
        while tail != None:
            if i == n:
                pre_node = head
            elif i > n:
                pre_node = pre_node.next
            tail = tail.next
            i += 1
        if pre_node == None:
            return head.next
        else:
            pre_node.next = pre_node.next.next
            return head
# @lc code=end
