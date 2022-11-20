package leetcode

// 300. Longest Increasing Subsequence 
//
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
// A subsequence is a sequence that can be derived from an array by deleting some or no elements 
// without changing the order of the remaining elements. 
// For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
//
// Runtime: 76 ms, faster than 40.00% of Go online submissions for Longest Increasing Subsequence.
// Memory Usage: 3.7 MB, less than 39.57% of Go online submissions for Longest Increasing Subsequence.
//
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = 1
    maxLen := 1
    for i := 1; i < len(nums); i++ {
        dp[i] = 1
        for j := i-1; j >= 0; j-- {
            if nums[i] > nums[j] && dp[j] >= dp[i]  {
                dp[i] = dp[j]+1
            }
        }
        if dp[i] > maxLen {
            maxLen = dp[i]
        }
    }
    return maxLen
}
