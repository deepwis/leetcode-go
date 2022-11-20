package leetcode

import (
    "math/rand"
)

// 215. Kth Largest Element in an Array
//
// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// You must solve it in O(n) time complexity.
//
// Runtime: 158 ms, faster than 57.96% of Go online submissions for Kth Largest Element in an Array.
// Memory Usage: 9.1 MB, less than 57.06% of Go online submissions for Kth Largest Element in an Array.
//
func findKthLargest(nums []int, k int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    p := rand.Intn(len(nums))
    index, pivot, high := -1, nums[p], len(nums) - 1
    nums[p], nums[high] = nums[high], nums[p] // move pivot to last, we will move it to medium position later.
    index = -1 // last left value (nums[index] >= pivot) index
    for i := 0; i < high; i++ {
        if nums[i] >= pivot {
            index++
            nums[i], nums[index] = nums[index], nums[i]
        }
    }
    index++
    nums[index], nums[high] = nums[high], nums[index] // move pivot back
    if index < k - 1 {
        return findKthLargest(nums[index+1:], k-index-1)
    } else if index > k - 1 {
        return findKthLargest(nums[:index], k)
    }
    return pivot
}
