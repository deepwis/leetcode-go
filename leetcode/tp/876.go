package leetcode

// 876. Middle of the Linked List
//
// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
//
// Runtime: 3 ms, faster than 36.50% of Go online submissions for Middle of the Linked List.
// Memory Usage: 2 MB, less than 15.37% of Go online submissions for Middle of the Linked List.
//
func middleNode(head *ListNode) *ListNode {
    middle, cur, i := head, head, 0
    for cur != nil {
        if i % 2 == 1{
            middle = middle.Next
        }
        cur = cur.Next
        i++
    }
    return middle
}
