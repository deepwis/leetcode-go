package leetcode

// 189. Rotate Array
//
// Given an array, rotate the array to the right by k steps, where k is non-negative.
//
// Runtime: 69 ms, faster than 43.14% of Go online submissions for Rotate Array.
// Memory Usage: 8.6 MB, less than 50.10% of Go online submissions for Rotate Array.
//
func rotate(nums []int, k int) {
    k = k % len(nums)
    if k == 0 {
        return
    }

    rotated := make([]int, k)
    copy(rotated, nums[len(nums)-k:])
    copy(nums[k:], nums[:len(nums)-k])
    copy(nums[:k], rotated)
}

