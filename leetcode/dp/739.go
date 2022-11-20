package leetcode

// 739. Daily Temperatures
//
// Given an array of integers temperatures represents the daily temperatures, return an array 
// answer such that answer[i] is the number of days you have to wait after the ith day to get 
// a warmer temperature. 
// If there is no future day for which this is possible, keep answer[i] == 0 instead.
//
// Runtime: 372 ms, faster than 29.05% of Go online submissions for Daily Temperatures.
// Memory Usage: 10.1 MB, less than 34.76% of Go online submissions for Daily Temperatures.
//
func dailyTemperatures(temperatures []int) []int {
    if len(temperatures) == 1 {
        return []int{0}
    }
    answer := make([]int, len(temperatures))
    q := []int{}
    for i := 0; i < len(temperatures); i++ {
        t := temperatures[i]
        q = append(q, i)
        j := len(q)-2
        for j >= 0 && temperatures[q[j]] < t {
            answer[q[j]] = i - q[j]
            q[j], q[j+1] = q[j+1], -1
            q = q[:j+1]
            j--
        }
    }
    return answer
}
