package leetcode

// 82. Remove Duplicates from Sorted List II
//
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, 
// leaving only distinct numbers from the original list. Return the linked list sorted as well.
//
// Runtime: 2 ms, faster than 96.02% of Go online submissions for Remove Duplicates from Sorted List II.
// Memory Usage: 2.9 MB, less than 87.17% of Go online submissions for Remove Duplicates from Sorted List II.
//
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    distinct := new(ListNode)
    distinct.Val = 101 // max node value is 100
    tail, prev, cur := distinct, distinct, head
    for cur != nil {
        next := cur.Next
        if prev.Val != cur.Val {
            if next == nil || next.Val != cur.Val {
                tail.Next = cur
                tail = cur
            }
        }
        prev = cur
        cur = next
    }
    tail.Next = nil
    return distinct.Next
}
