package leetcode

// 160. Intersection of Two Linked Lists
//
// Runtime: 31 ms, faster than 95.31% of Go online submissions for Intersection of Two Linked Lists.
// Memory Usage: 6.8 MB, less than 92.46% of Go online submissions for Intersection of Two Linked Lists.
//
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1, p2 := headA, headB
    for p1 != nil || p2 != nil{
        if p1 == p2 {
            return p1
        }
        if p1 == nil {
            p1 = headB
        } else {
            p1 = p1.Next
        }
        if p2 == nil {
            p2 = headA
        } else {
            p2 = p2.Next
        }
    }
    return nil
}

