package leetcode

// 143. Reorder List
//
// You are given the head of a singly linked-list. The list can be represented as:
//   L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
//   L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.
//
// Runtime: 7 ms, faster than 97.64% of Go online submissions for Reorder List.
// Memory Usage: 5.3 MB, less than 100.00% of Go online submissions for Reorder List.
//
func reorderList(head *ListNode) {
    var reversed *ListNode
    if head.Next == nil {
        return
    }
    slow := new(ListNode) // We define a slow pinter to prev of head node
    slow.Next = head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // Find tail of left part
    leftTail := slow
    if fast != nil {
        leftTail = slow.Next
    }
    // We reverse the right part
    slow = leftTail.Next
    leftTail.Next = nil
    for slow != nil {
        next := slow.Next
        slow.Next = reversed
        reversed = slow
        slow = next
    }
    for reversed != nil {
        next := reversed.Next
        reversed.Next = head.Next
        head.Next = reversed
        head = reversed.Next
        reversed = next
    }
}
