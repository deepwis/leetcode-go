package leetcode

// 153. Find Minimum in Rotated Sorted Array
//
// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. 
// For example, the array nums = [0,1,2,4,5,6,7] might become:
// * [4,5,6,7,0,1,2] if it was rotated 4 times.
// * [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results 
// in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
// You must write an algorithm that runs in O(log n) time.
//
// Runtime: 4 ms, faster than 59.71% of Go online submissions for Find Minimum in Rotated Sorted Array.
// Memory Usage: 2.5 MB, less than 80.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
//
func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    low, high := 0, len(nums)
    for low <= high {
        if nums[low] < nums[high-1] {
            return nums[low]
        }
        mid := (low + high) / 2
        if mid == high - 1 {
            return nums[high-1]
        }
        if nums[low] <= nums[mid] {
            low = mid+1
        } else {
            high = mid+1
        }
    }
    return nums[low]
}

