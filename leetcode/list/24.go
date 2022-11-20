package leetcode

// 24. Swap Nodes in Pairs
//
// Given a linked list, swap every two adjacent nodes and return its head. 
// You must solve the problem without modifying the values in the list's nodes 
// (i.e., only nodes themselves may be changed.)
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
// Memory Usage: 2.1 MB, less than 78.46% of Go online submissions for Swap Nodes in Pairs.
//
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }

    cur := head
    swapped := head.Next
    cur.Next = swapped.Next
    swapped.Next = cur

    for cur.Next != nil {
        n1 := cur.Next
        n2 := n1.Next
        if n2 == nil {
            break
        }
        n1.Next = n2.Next
        n2.Next = n1
        cur.Next = n2
        cur = n1
    }
    return swapped
}

