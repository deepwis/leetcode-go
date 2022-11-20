package leetcode

// 160. Intersection of Two Linked Lists
//
// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. 
// If the two linked lists have no intersection at all, return null.
//
// Runtime: 85 ms, faster than 18.85% of Go online submissions for Intersection of Two Linked Lists.
// Memory Usage: 8.1 MB, less than 13.13% of Go online submissions for Intersection of Two Linked Lists.
//
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    visited := make(map[*ListNode]bool)
    cur := headA
    for cur != nil {
        visited[cur] = true
        cur = cur.Next
    }

    cur = headB
    for cur != nil {
        if _, ok := visited[cur]; ok {
            return cur
        }
        cur = cur.Next
    }
    return nil
}
