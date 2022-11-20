package leetcode

// 75. Sort Colors
//
// Given an array nums with n objects colored red, white, or blue, sort them in-place so that 
// objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.
//
// Runtime: 5 ms, faster than 10.42% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 71.07% of Go online submissions for Sort Colors.
//
func sortColors(nums []int) {
    first1, last1 := 0, len(nums) - 1
    for first1 < last1 {
        if nums[first1] == 0 {
            first1++
        } else if nums[first1] == 2 {
            nums[first1], nums[last1] = nums[last1], nums[first1]
            last1--
        } else if nums[last1] == 0 {
            nums[first1], nums[last1] = nums[last1], nums[first1]
            first1++
        } else if nums[last1] == 2 {
            last1--
        } else {
            break
        }
    }
    for i := first1+1; i <= last1; {
        if nums[i] == 0 {
            nums[first1], nums[i] = 0, 1
            first1++
            i++
        } else if nums[i] == 2 {
            nums[last1], nums[i] = 2, nums[last1]
            last1--
        } else {
            i++
        }
    }
}
