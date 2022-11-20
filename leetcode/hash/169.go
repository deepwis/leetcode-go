package leetcode

// 169. Majority Element
//
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. 
// You may assume that the majority element always exists in the array.
//
// Runtime: 46 ms, faster than 16.36% of Go online submissions for Majority Element.
// Memory Usage: 6.4 MB, less than 10.34% of Go online submissions for Majority Element.
//
func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    elemCnt := make(map[int]int, (len(nums) + 1) / 2)
    for _, i := range nums {
        elemCnt[i] += 1
        if elemCnt[i] > len(nums) / 2 {
            return i
        }
    }
    return nums[0]
}
