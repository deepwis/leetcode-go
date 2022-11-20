package leetcode

// 287. Find the Duplicate Number
//
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.
//
// Runtime: 100 ms, faster than 88.22% of Go online submissions for Find the Duplicate Number.
// Memory Usage: 8.9 MB, less than 62.36% of Go online submissions for Find the Duplicate Number.
//
func findDuplicate(nums []int) int {
    for i := 0; i < len(nums) - 1; {
        if nums[i] != i {
            k := nums[i]
            if nums[k] == nums[i] {
                return nums[i]
            }
            nums[i], nums[k] = nums[k], nums[i]
        } else {
            i++
        }
    }
    return nums[len(nums)-1]
}
