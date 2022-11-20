package leetcode

// 142. Linked List Cycle II
//
// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again 
// by continuously following the next pointer. 
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). 
// It is -1 if there is no cycle. Note that pos is not passed as a parameter.
// Do not modify the linked list.
//
// Runtime: 8 ms, faster than 76.58% of Go online submissions for Linked List Cycle II.
// Memory Usage: 5.2 MB, less than 20.99% of Go online submissions for Linked List Cycle II.
//
func detectCycle(head *ListNode) *ListNode {
    visited := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if _, ok := visited[cur]; ok {
            return cur
        } else {
            visited[cur] = true
        }
        cur = cur.Next
    }
    return nil
}
