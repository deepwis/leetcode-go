package leetcode

// 162. Find Peak Element
//
// A peak element is an element that is strictly greater than its neighbors.
// Given a 0-indexed integer array nums, find a peak element, and return its index. 
// If the array contains multiple peaks, return the index to any of the peaks.
// You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to 
// be strictly greater than a neighbor that is outside the array.
// You must write an algorithm that runs in O(log n) time.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Peak Element.
// Memory Usage: 2.8 MB, less than 71.48% of Go online submissions for Find Peak Element.
//
func findPeakElement(nums []int) int {
    n := len(nums)
    low, high := 0, n-1
    for low < high-1 {
        mid := low + (high - low) >> 1
        if mid == low {
            break
        }
        if nums[mid] < nums[mid+1] {
            low = mid
        } else if nums[mid] < nums[mid-1] {
            high = mid
        } else {
            return mid
        }
    }
    if nums[low] > nums[high] {
        return low
    }
    return high
}

// findMax() find Max Peak Element
//
// Runtime: 8 ms, faster than 34.71% of Go online submissions for Find Peak Element.
// Memory Usage: 2.8 MB, less than 71.48% of Go online submissions for Find Peak Element.
//
func findMax(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    if n == 2 {
        if nums[0] > nums[1] {
            return 0
        } else {
            return 1
        }
    }
    mid := n / 2
    l := findMax(nums[:mid])
    r := findMax(nums[mid:]) + mid
    if nums[l] > nums[r] {
        return l
    }
    return r
}
