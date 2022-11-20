package leetcode

// 560. Subarray Sum Equals K
//
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Runtime: 1728 ms, faster than 11.27% of Go online submissions for Subarray Sum Equals K.
// Memory Usage: 7 MB, less than 79.72% of Go online submissions for Subarray Sum Equals K.
//
func subarraySum(nums []int, k int) int {
    if len(nums) == 1 {
        if nums[0] == k {
            return 1
        }
        return 0
    }
    total := 0
    sums := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        for j := 0; j <= i; j++ {
            sums[j] = sums[j] + nums[i]
            if sums[j] == k {
                total++
            }
        }
    }
    return total
}
