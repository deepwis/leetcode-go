package leetcode

import "math/rand"

// 912. Sort an Array
//
// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity 
// and with the smallest space complexity possible.
//
// Runtime: 860 ms, faster than 37.28% of Go online submissions for Sort an Array.
// Memory Usage: 11.9 MB, less than 32.73% of Go online submissions for Sort an Array.
//
func sortArray(nums []int) []int {
    sortRange(nums, 0, len(nums)-1)
    return nums
}

func sortRange(nums []int, s, e int) {
    if e - s < 1 {
        return
    }
    p := rand.Intn(e+1-s) + s
    pivot := nums[p]
    nums[p], nums[e] = nums[e], nums[p]
    idx := s-1
    for i := s; i < e; i++ {
        if nums[i] < pivot {
            idx++
            nums[i], nums[idx] = nums[idx], nums[i]
        }
    }
    nums[idx+1], nums[e] = nums[e], nums[idx+1]
    sortRange(nums, s, idx)
    sortRange(nums, idx+2, e)
}

