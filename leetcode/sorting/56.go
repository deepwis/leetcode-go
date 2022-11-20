package leetcode

import (
    "sort"
)

// 56. Merge Intervals
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, 
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Runtime: 50 ms, faster than 29.74% of Go online submissions for Merge Intervals.
// Memory Usage: 7.4 MB, less than 13.83% of Go online submissions for Merge Intervals.
//
func merge(intervals [][]int) [][]int{
    if len(intervals) == 1 {
        return intervals
    }
    sort.Sort(Intervals(intervals))
    merged := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        k := len(merged) - 1
        if intervals[i][0] - merged[k][1] < 1 {
            end := intervals[i][1]
            if intervals[i][1] < merged[k][1] {
                end = merged[k][1]
            }
            merged[k] = []int{merged[k][0], end}
        } else {
            merged = append(merged, intervals[i])
        }
    }
    return merged
}

type Intervals [][]int
func (x Intervals) Len() int { return len(x) }
func (x Intervals) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x Intervals) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
