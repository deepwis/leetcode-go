package leetcode

// 206. Reverse Linked List
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Runtime: 5 ms, faster than 36.47% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 79.13% of Go online submissions for Reverse Linked List.
// 
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var prev, next *ListNode

    cur := head
    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev, cur = cur, next
    }
    return prev
}

