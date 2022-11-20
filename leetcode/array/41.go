package leetcode

// 41. First Missing Positive
//
// Given an unsorted integer array nums, return the smallest missing positive integer.
// You must implement an algorithm that runs in O(n) time and uses constant extra space.
//
// Runtime: 69 ms, faster than 59.50% of Go online submissions for First Missing Positive.
// Memory Usage: 9 MB, less than 32.25% of Go online submissions for First Missing Positive.
//
func firstMissingPositive(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 1
    }

    i := 0
    for i < n {
        k := nums[i] - 1
        if nums[i] <=0 || nums[i] > n || nums[i] == i+1 || nums[i] == nums[k] {
            i++
            continue
        }
        nums[i], nums[k] = nums[k], nums[i]
    }
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return n+1
}
