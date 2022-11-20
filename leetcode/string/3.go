package leetcode

// 3. Longest Substring Without Repeating Characters
//
// Given a string s, find the length of the longest substring without repeating characters.
//
// Runtime: 8 ms, faster than 82.87% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 3.2 MB, less than 49.07% of Go online submissions for Longest Substring Without Repeating 
//
func lengthOfLongestSubstring(s string) int {
    var maxLen int
    visited := make(map[rune]int)
    start := 0
    for i, c := range s {
        idx, ok := visited[c]
        if ok && idx >= start {
            if i - start > maxLen {
                maxLen = i - start
            }
            start = idx + 1
        }
        visited[c] = i
    }
    if len(s) - start > maxLen {
        maxLen = len(s) - start
    }

    return maxLen
}

