package leetcode

// 53. Maximum Subarray
//
// Given an integer array nums, find the contiguous subarray (containing at least one number) 
// which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.
//
// Runtime: 236 ms, faster than 14.86% of Go online submissions for Maximum Subarray.
// Memory Usage: 9.3 MB, less than 70.21% of Go online submissions for Maximum Subarray.
//
func maxSubArray(nums []int) int {
    max := nums[0]
    s, n := -1, len(nums)
    for i := 0;  i < n; i++ {
        if nums[i] > 0 {
            s = i
            break
        }
        if max < nums[i] {
            max = nums[i]
        }
    }
    if s == -1 {
        return max
    }
    cur := 0
    for i := s; i < n; i++ {
        cur += nums[i]
        if nums[i] > 0 && cur > max {
            max = cur
        }
        if cur < 0 {
            cur = 0
        }
    }
    return max
}
