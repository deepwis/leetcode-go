package leetcode

// 141. Linked List Cycle
//
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached again 
// by continuously following the next pointer. 
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to. 
// Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// Runtime: 20 ms, faster than 22.46% of Go online submissions for Linked List Cycle.
// Memory Usage: 6.4 MB, less than 17.18% of Go online submissions for Linked List Cycle.
//
func HasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if _, ok := visited[cur]; ok {
            return true
        }
        visited[cur] = true
        cur = cur.Next
    }
    return false
}
