package leetcode

// 647. Palindromic Substrings
//
// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome when it reads the same backward as forward.
// A substring is a contiguous sequence of characters within the string.
//
// Runtime: 90 ms, faster than 11.90% of Go online submissions for Palindromic Substrings.
// Memory Usage: 1.9 MB, less than 94.64% of Go online submissions for Palindromic Substrings.
//
func countSubstrings(s string) int {
    if len(s) <= 1 {
        return len(s)
    }
    cnt := 0
    for i := 0; i < len(s); i++ {
        cnt++
        for j := 0; j<i; j++ {
            if isPalindromeString(s[j:i+1]) {
                cnt++
            }
        }
    }
    return cnt
}
