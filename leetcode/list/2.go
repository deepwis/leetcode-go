package leetcode

// 2. Add Two Numbers
//
// You are given two non-empty linked lists representing two non-negative integers. 
// The digits are stored in reverse order, and each of their nodes contains a single digit. 
// Add the two numbers and return the sum as a linked list.
//
// Runtime: 22 ms, faster than 32.84% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.6 MB, less than 47.70% of Go online submissions for Add Two Numbers.
//
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    v1, v2 := l1.Val, l2.Val
    v, f := (v1 + v2) % 10, (v1 + v2) / 10
    sum := &ListNode{v, nil, nil}
    cur1, cur2, curS := l1.Next, l2.Next, sum
    for cur1 != nil || cur2 != nil {
        if cur1 != nil {
            v1 = cur1.Val
            cur1 = cur1.Next
        } else {
            v1 = 0
        }
        if cur2 != nil {
            v2 = cur2.Val
            cur2 = cur2.Next
        } else {
            v2 = 0
        }
        v, f = (v1 + v2 + f) % 10, (v1 + v2 + f) / 10
        n := &ListNode{v, nil, nil}
        curS.Next = n
        curS = n
    }
    if f != 0 {
        n := &ListNode{f, nil, nil}
        curS.Next = n
    }
    return sum
}
