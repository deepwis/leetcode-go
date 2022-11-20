package leetcode

// 704. Binary Search
//
// Given an array of integers nums which is sorted in ascending order, and an integer target, 
// write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
// You must write an algorithm with O(log n) runtime complexity.
//
// Runtime: 24 ms, faster than 99.68% of Go online submissions for Binary Search.
// Memory Usage: 6.7 MB, less than 75.19% of Go online submissions for Binary Search.
//
func search2(nums []int, target int) int {
    low, high := 0, len(nums)
    for low < high {
        mid := low + (high - low) >> 1
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            high = mid
        } else {
            low = mid+1
        }
    }
    return -1
}
