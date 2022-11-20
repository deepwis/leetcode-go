package leetcode

// 718. Maximum Length of Repeated Subarray
//
// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.
//
// Runtime: 72 ms, faster than 64.41% of Go online submissions for Maximum Length of Repeated Subarray.
// Memory Usage: 19.1 MB, less than 43.83% of Go online submissions for Maximum Length of Repeated Subarray.
//
func findLength(nums1, nums2 []int) int {
    m, n, max := len(nums1), len(nums2), 0
    dp := make([][]int, m+1)
    dp[0] = make([]int, n+1)
    for i := 0; i < m; i++ {
        dp[i+1] = make([]int, n+1)
        for j := 0; j < n; j++ {
            if nums1[i] == nums2[j] {
                dp[i+1][j+1] = dp[i][j] + 1
                if max < dp[i+1][j+1] {
                    max = dp[i+1][j+1]
                }
            }
        }
    }
    return max
}
