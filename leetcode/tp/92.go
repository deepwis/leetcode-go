package leetcode

// 92. Reverse Linked List II
//
// Given the head of a singly linked list and two integers left and right where left <= right, 
// reverse the nodes of the list from position left to position right, and return the reversed list.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 16.51% of Go online submissions for Reverse Linked List II.
//
func reverseBetween(head *ListNode, left, right int) *ListNode {
    if left == 1 && right == 1 {
        return head
    }
    slow ,fast := head, head
    i := 1
    for fast != nil && fast.Next != nil && i < right {
        if i == left-1 {
            slow = fast
        } else if i == left-2 {
            slow = fast.Next
        }
        if i == right {
            fast = fast
            break
        } else if i == right-1 {
            fast = fast.Next
            break
        }
        i += 2
        fast = fast.Next.Next
    }
    cur, r := head, fast.Next
    if left > 1 {
        cur = slow.Next
    }
    fast.Next = nil
    for cur != nil {
        next := cur.Next
        cur.Next = r
        r = cur
        cur = next
    }
    if left == 1 {
        return r
    }
    slow.Next = r
    return head
}
