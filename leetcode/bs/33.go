package leetcode

// 33. Search in Rotated Sorted Array
//
// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index 
// k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], 
// nums[0], nums[1], ..., nums[k-1]] (0-indexed). 
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
// Given the array nums after the possible rotation and an integer target, return the index of target 
// if it is in nums, or -1 if it is not in nums.
// You must write an algorithm with O(log n) runtime complexity.
//
// Runtime: 4 ms, faster than 65.36% of Go online submissions for Search in Rotated Sorted Array.
// Memory Usage: 2.6 MB, less than 72.33% of Go online submissions for Search in Rotated Sorted Array.
//
func search(nums []int, target int) int {
    if len(nums) == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
    l, r := 0, len(nums) - 1
    for l < r {
        if target == nums[l] {
            return l
        }
        if target == nums[r] {
            return r
        }
        if r - l == 1 {
            return -1
        }
        m := (l + r) / 2
        if target == nums[m] {
            return m
        }
        if nums[m] < nums[l] { // l < 0 < m
            if target > nums[m] && target < nums[r] { // m < target < r
                l = m + 1
            } else { // l < target < m
                r = m - 1
            }
        } else { // m < 0 < r
            if target < nums[m] && target > nums[l] {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }
    return -1
}

