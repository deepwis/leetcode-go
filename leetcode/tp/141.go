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
// Runtime: 9 ms, faster than 87.26% of Go online submissions for Linked List Cycle.
// Memory Usage: 4.4 MB, less than 31.25% of Go online submissions for Linked List Cycle.
//
func HasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        if fast == slow  || fast.Next == slow {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return false
}
