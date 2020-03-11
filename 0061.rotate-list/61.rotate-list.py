#
# @lc app=leetcode id=61 lang=python3
#
# [61] Rotate List
#

# @lc code=start
# Definition for singly-linked list.


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if head == None or head.next == None or k == 0:
            return head
        n = self.length(head)
        k = k % n
        if k == 0:
            return head
        tail, nodes = head, []
        while tail != None:
            nodes.append(tail)
            tail = tail.next
        tail = nodes.pop()
        for _ in range(k):
            last = nodes.pop()
            last.next = None
            tail.next = head
            head = tail
            tail = last
        return head

    def length(self, head: ListNode) -> int:
        n, node = 1, head
        while node.next != None:
            n += 1
            node = node.next
        return n

# @lc code=end
