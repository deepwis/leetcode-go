package leetcode

// 19. Remove Nth Node from End of List
//
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.1 MB, less than 87.27% of Go online submissions for Remove Nth Node From End of List.
//
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {
        return head
    }
    var prev, nth *ListNode
    i, cur := 0, head
    for cur != nil {
        i++
        if i == n {
            nth = head
        }
        if i > n {
            prev = nth
            nth = nth.Next
        }
        cur = cur.Next
    }
    if i < n {
        return head
    }
    if prev == nil {
        return nth.Next
    }
    prev.Next = nth.Next
    return head
}

