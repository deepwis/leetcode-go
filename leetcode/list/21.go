package leetcode

// 21. Merge Two Sorted Lists
//
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists in a one sorted list. 
// The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.6 MB, less than 50.16% of Go online submissions for Merge Two Sorted Lists.
//
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    var ret, cur, next, other *ListNode
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    if list1.Val <= list2.Val {
        cur, other = list1, list2
    } else {
        cur, other = list2, list1
    }
    ret = cur
    for cur != nil && other != nil {
        next = cur.Next
        if next != nil && next.Val <= other.Val {
            cur = next
        } else {
            cur.Next = other
            cur, other = other, next
        }
    }

    return ret
}

