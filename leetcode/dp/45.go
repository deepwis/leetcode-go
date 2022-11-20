package leetcode

// 45. Jump Game II
//
// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
// Each element nums[i] represents the maximum length of a forward jump from index i. 
// In other words, if you are at nums[i], you can jump to any nums[i + j] where:
// * 0 <= j <= nums[i] and
// * i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. 
// The test cases are generated such that you can reach nums[n - 1].
//
// Runtime: 22 ms, faster than 83.19% of Go online submissions for Jump Game II.
// Memory Usage: 6 MB, less than 92.02% of Go online submissions for Jump Game II.
//
const MAXJUMP = 99999

func jump(nums []int) int {
    ret, n := 0, len(nums)
    minJumps := []int{n-1} // index is min step, elements is position
    for i := n-2; i >= 0; i-- {
        maxTo := i + nums[i]
        if maxTo == i {
            continue
        }
        for k := 0; k < len(minJumps); k++ {
            ret = MAXJUMP
            if maxTo >= minJumps[k] {
                ret = k+1
                if k == len(minJumps)-1 {
                    minJumps = append(minJumps, i)
                } else {
                    minJumps[k+1] = i
                }
                break
            }
        }
    }
    return ret
}
