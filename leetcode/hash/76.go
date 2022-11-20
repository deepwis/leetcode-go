package leetcode

// 76. Minimum Window Substring
//
// Given two strings s and t of lengths m and n respectively, return the minimum window substring of s 
// such that every character in t (including duplicates) is included in the window. 
// If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.
// A substring is a contiguous sequence of characters within the string.
//
// Runtime: 17 ms, faster than 78.20% of Go online submissions for Minimum Window Substring.
// Memory Usage: 2.9 MB, less than 24.94% of Go online submissions for Minimum Window Substring.
//
func minWindow(s, t string) string {
    if len(s) < len(t) {
        return ""
    }
    sub := ""
    needs := len(t)
    charCnt := make(map[byte]int, needs)
    for i := 0; i < needs; i++ {
        if _, ok := charCnt[t[i]]; ok {
            charCnt[t[i]] -= 1
        } else {
            charCnt[t[i]] = 0
        }
    }
    low, high := 0, 0
    for high < len(s) {
        if needs > 0 {
            if c, ok := charCnt[s[high]]; ok {
                if c <= 0 {
                    needs--
                }
                charCnt[s[high]] += 1
            }
            high++
        }
        for needs <= 0 {
            if c, ok := charCnt[s[low]]; ok {
                if len(sub) == 0 || high - low < len(sub) {
                    sub = s[low:high]
                }
                if c <= 1 {
                    needs++
                }
                charCnt[s[low]] -= 1
            }
            low++
        }
    }
    return sub
}
