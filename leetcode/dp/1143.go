package leetcode

// 1143. Longest Common Subsequence
//
// Given two strings text1 and text2, return the length of their longest common subsequence. 
// If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some characters 
// (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.
//
// Runtime: 7 ms, faster than 70.03% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 11 MB, less than 36.50% of Go online submissions for Longest Common Subsequence.
//
func longestCommonSubsequence(text1, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    dp[0] = make([]int, n+1)
    for i := 1; i <= len(text1); i++ {
        dp[i] = make([]int, n+1)
        for j := 1; j <= len(text2); j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1]+1
            }
            if dp[i][j] < dp[i-1][j] {
                dp[i][j] = dp[i-1][j]
            }
            if dp[i][j] < dp[i][j-1] {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[m][n]
}
