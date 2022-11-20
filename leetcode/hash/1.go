package leetcode

// 1. Two Sum
//
// Given an array of integers nums and an integer target, return indices of 
// the two numbers such that they add up to target.
//
// Runtime: 17 ms, faster than 48.16% of Go online submissions for Two Sum.
// Memory Usage: 4.4 MB, less than 39.11% of Go online submissions for Two Sum.
//
func twoSum(nums []int, target int) []int {
    ret := make([]int, 2, 2)
    numIndex := make(map[int]int)
    for i, v := range nums {
        other := target - v
        j, ok := numIndex[other]
        if ok {
            ret[0] = j
            ret[1] = i
            break
        }
        numIndex[v] = i
    }
    return ret
}
