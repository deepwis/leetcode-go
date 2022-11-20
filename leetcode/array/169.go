package leetcode

// 169. Majority Element
//
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. 
// You may assume that the majority element always exists in the array.
//
// Runtime: 25 ms, faster than 73.30% of Go online submissions for Majority Element.
// Memory Usage: 6.3 MB, less than 34.77% of Go online submissions for Majority Element.
// 
func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    candidate, c := nums[0], 1
    for _, i := range nums[1:] {
        if c == 0 {
            candidate = i
        }
        if candidate == i {
            c++
        } else {
            c--
        }
    }
    return candidate
}
