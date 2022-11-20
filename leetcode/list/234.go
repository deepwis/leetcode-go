package leetcode

// 234. Palindrome Linked List
//
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
//
// Runtime: 357 ms, faster than 27.75% of Go online submissions for Palindrome Linked List.
// Memory Usage: 12.7 MB, less than 17.25% of Go online submissions for Palindrome Linked List.
//
func isPalindrome(head *ListNode) bool {
    var valArr []int
    cur := head
    for cur != nil {
        valArr = append(valArr, cur.Val)
        cur = cur.Next
    }
    for i, j := 0, len(valArr) - 1; i < j; {
        if valArr[i] != valArr[j] {
            return false
        }
        i++
        j--
    }
    return true
}

