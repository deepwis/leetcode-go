package leetcode

// 209. Minimum Size Subarray Sum
//
// Given an array of positive integers nums and a positive integer target, return the minimal length of 
// a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
//
// Runtime: 41 ms, faster than 80.32% of Go online submissions for Minimum Size Subarray Sum.
// Memory Usage: 8.4 MB, less than 26.60% of Go online submissions for Minimum Size Subarray Sum.
//
func miniSubArrayLen(target int, nums []int) int {
    n := len(nums)
    sum, ans := 0, n+1
    low, high := 0, 0
    for high <= n {
        if sum >= target {
            if high - low < ans {
                ans = high - low
            }
            sum -= nums[low]
            low++
            continue
        }
        if high == n {
            break
        }
        sum += nums[high]
        high++
    }
    if ans > n {
        return 0
    }
    return ans
}

