package leetcode

// 5. Longest Palindromic Substring
//
// Given a string s, return the longest palindromic substring in s.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.4 MB, less than 88.40% of Go online submissions for Longest Palindromic Substring.
//
func longestPalindrome(s string) string {
    var ret string
    for i := 0; i < len(s); i++ {
        s1 := expandAroundCenter(s, i, i)
        s2 := expandAroundCenter(s, i, i + 1)
        if len(s1) > len(ret) {
            ret = s1
        }
        if len(s2) > len(ret) {
            ret = s2
        }
    }
    return ret
}

func expandAroundCenter(s string, left, right int) string {
    if left < 0 || right >= len(s) || s[left] != s[right] {
        return ""
    }
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left+1:right]
}


