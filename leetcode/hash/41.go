package leetcode

// 41. First Missing Positive
//
// Given an unsorted integer array nums, return the smallest missing positive integer.
// You must implement an algorithm that runs in O(n) time and uses constant extra space.
//
// Runtime: 63 ms, faster than 64.50% of Go online submissions for First Missing Positive.
// Memory Usage: 8.3 MB, less than 57.75% of Go online submissions for First Missing Positive.
//
func firstMissingPositive(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 1
    }

    l, r := 0, n
    for l < r {
        k := nums[l] - 1
        if nums[l] == l+1 {
            // number is in right position
            l++
        } else if nums[l] <= l || nums[l] > r || nums[l] == nums[k] {
            // found extra numbers, wap to right, we don't need it
            nums[l], nums[r-1] = nums[r-1], nums[l]
            r--
        } else  {
            // number is not in right position, swap to right position.
            nums[l], nums[k] = nums[k], nums[l]
        }
    }
    return l+1
}
