package leetcode

// 25. Reverse Nodes in k-Group
//
// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
// k is a positive integer and is less than or equal to the length of the linked list. 
// If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// Runtime: 6 ms, faster than 76.96% of Go online submissions for Reverse Nodes in k-Group.
// Memory Usage: 3.8 MB, less than 11.16% of Go online submissions for Reverse Nodes in k-Group.
//
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k == 1 {
        return head
    }

    var reversed, tail *ListNode

    stack, ng := make([]*ListNode, k), 0
    cur := head
    for cur != nil {
        stack[ng] = cur
        cur = cur.Next
        ng++
        if ng == k {
            if reversed == nil {
                ng--
                reversed = stack[ng]
                tail = reversed
                stack[ng] = nil
            }
            for ng > 0 {
                ng--
                tail.Next = stack[ng]
                tail = stack[ng]
                stack[ng] = nil
            }
        }
    }
    if tail != nil {
        tail.Next = nil
    }

    if ng > 0 {
        if reversed == nil {
            reversed = stack[0]
        } else {
            tail.Next = stack[0]
        }
    }
    return reversed
}

