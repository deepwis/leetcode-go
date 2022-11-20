package leetcode

// 136. Single Number
//
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
// Runtime: 25 ms, faster than 74.01% of Go online submissions for Single Number.
// Memory Usage: 6.6 MB, less than 35.49% of Go online submissions for Single Number.
//
func singleNumber(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    ret := nums[0]
    for i := 1; i < len(nums); i++ {
        ret ^= nums[i]
    }
    return ret
}
