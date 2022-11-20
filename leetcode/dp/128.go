package leetcode

// 128. Longest Consecutive Sequence
//
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
//
// Runtime: 152 ms, faster than 62.71% of Go online submissions for Longest Consecutive Sequence.
// Memory Usage: 11.4 MB, less than 37.10% of Go online submissions for Longest Consecutive Sequence.
func longestConsecutive(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    dp := make(map[int]int, len(nums))
    maxLen := 0

    for _, i := range nums {
        if _, ok := dp[i]; !ok {
            dp[i] = 1
            left, lok := dp[i-1]
            right, rok := dp[i+1]
            dp[i] = left + right + 1
            if lok {
                dp[i-left] = dp[i] // update left bound [i-left, i]
            }
            if rok {
                dp[i+right] = dp[i] // // update right bound [i, i+right]
            }
            if dp[i] > maxLen {
                maxLen = dp[i]
            }
        }
    }
    return maxLen
}

