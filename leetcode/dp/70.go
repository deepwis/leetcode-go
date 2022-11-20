package leetcode

// 70. Climbing Stairs
//
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 2 MB, less than 39.58% of Go online submissions for Climbing Stairs.
//
func climbStairs(n int) int {
    if n == 0 {
        return 0
    }
    steps := make([]int, n) // steps for every stairs, steps for k stair is steps[k-1]
    steps[0] = 1
    if n == 1 {
        return steps[0]
    }
    steps[1] = 2
    for i := 2; i < n; i++ {
        // before we arrives at stairs i, we must arrives at stairs i-1 or i-2 at first.
        // for every step we arrives at step i-1, we have 1 ways to arrived at step i.
        // for every step we arrives at step i-2, we have 2 ways to arrived at step i. But only 1 way can skip step 1-1
        steps[i] = steps[i-2] + steps[i-1]
    }
    return steps[n-1]
}
