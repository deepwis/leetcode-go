package leetcode

// 239. Sliding Window Maximum
//
// You are given an array of integers nums, there is a sliding window of size k which is moving from 
// the very left of the array to the very right. You can only see the k numbers in the window. 
// Each time the sliding window moves right by one position.
// Return the max sliding window.
//
// Runtime: 469 ms, faster than 62.85% of Go online submissions for Sliding Window Maximum.
// Memory Usage: 10 MB, less than 80.41% of Go online submissions for Sliding Window Maximum.
//
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 1 {
        return []int{nums[0]}
    }
    if k == 1 {
        return nums
    }
    maximums, left := make([]int, len(nums)-k+1), -k+1
    dq := make([]int, 0, k)
    dq = append(dq, 0)
    for i := 1; i < len(nums); i++ {
        left++
        for len(dq) > 0 {
            tail := dq[len(dq)-1]
            if nums[tail] < nums[i] {
                dq = dq[:len(dq)-1]
            } else {
                break
            }
        }
        dq = append(dq, i)
        if left >= 0 {
            maximums[left] = nums[dq[0]]
        }
        if len(dq) == k || left == dq[0] {
            dq = dq[1:]
        }
    }
    return maximums
}

