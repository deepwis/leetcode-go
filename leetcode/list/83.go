package leetcode

// 83. Remove Duplicates from Sorted List
//
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. 
// Return the linked list sorted as well.
//
// Runtime: 3 ms, faster than 88.19% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 41.45% of Go online submissions for Remove Duplicates from Sorted List.
//
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        if slow.Val != fast.Val {
            slow.Next = fast
            slow = fast
        }
        if slow.Val != fast.Next.Val {
            slow.Next = fast.Next
            slow = fast.Next
        }
        fast = fast.Next.Next
    }
    if fast != nil && slow.Val == fast.Val {
        slow.Next = nil
    } else {
        slow.Next = fast
    }
    return head
}
