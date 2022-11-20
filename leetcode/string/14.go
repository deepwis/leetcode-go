package leetcode

// 14. Longest Common Prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.4 MB, less than 54.69% of Go online submissions for Longest Common Prefix.
//
func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        j := 0
        for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
            j++
        }
        if j == 0 {
            return ""
        }
        prefix = prefix[:j]
    }
    return prefix
}
