package leetcode

// 55. Jump Game
//
// You are given an integer array nums. You are initially positioned at the array's first index, 
// and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.
//
// Runtime: 63 ms, faster than 91.15% of Go online submissions for Jump Game.
// Memory Usage: 6.9 MB, less than 87.59% of Go online submissions for Jump Game.
//
func canJump(nums []int) bool {
    allCan := []int{len(nums)-1}
    for i := len(nums) - 2; i >= 0; i-- {
        nextCan := allCan[len(allCan)-1]
        if i + nums[i] >= nextCan {
            allCan = append(allCan, i)
        }
    }
    return allCan[len(allCan)-1] == 0
}

// Runtime: 78 ms, faster than 85.44% of Go online submissions for Jump Game.
// Memory Usage: 7.4 MB, less than 65.80% of Go online submissions for Jump Game.
//
func canJump2(nums []int) bool {
    nextCan := len(nums)-1
    for i := len(nums) - 2; i >= 0; i-- {
        if i + nums[i] >= nextCan {
            nextCan = i
        }
    }
    return nextCan == 0
}
