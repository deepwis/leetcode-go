package leetcode

// 234. Palindrome Linked List
//
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
//
// Runtime: 209 ms, faster than 82.14% of Go online submissions for Palindrome Linked List.
// Memory Usage: 9.2 MB, less than 95.69% of Go online submissions for Palindrome Linked List.
//
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    var reversed *ListNode
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    for slow != nil {
        next := slow.Next
        slow.Next = reversed
        reversed = slow
        slow = next
    }

    for reversed != nil {
        if reversed.Val != head.Val {
            return false
        }
        reversed = reversed.Next
        head = head.Next
    }
    return true
}
