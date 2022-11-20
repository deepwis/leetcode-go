package leetcode

// 283. Move Zeroes
//
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order 
// of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
//
// Runtime: 39 ms, faster than 67.76% of Go online submissions for Move Zeroes.
// Memory Usage: 7.5 MB, less than 16.58% of Go online submissions for Move Zeroes.
//
func moveZeroes(nums []int) {
    if len(nums) < 2{
        return
    }
    for low, high := 0, 1; high < len(nums); high++ {
        if nums[low] != 0  {
            low++
            continue
        } else if nums[high] != 0 {
            nums[low], nums[high] = nums[high], nums[low]
            low++
        }
    }
}

