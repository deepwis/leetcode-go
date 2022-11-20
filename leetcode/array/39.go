package leetcode

import "sort"

// 39. Combination Sum
//
// Given an array of distinct integers candidates and a target integer target, return a list of 
//all unique combinations of candidates where the chosen numbers sum to target. 
// You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times. 
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
// The test cases are generated such that the number of unique combinations that sum up to target 
// is less than 150 combinations for the given input.
//
// Runtime: 8 ms, faster than 42.42% of Go online submissions for Combination Sum.
// Memory Usage: 3.2 MB, less than 60.32% of Go online submissions for Combination Sum.
//
func combinationSum(candidates []int, target int) [][]int{
    cIndex := make(map[int]int)
    targetSets := [][]int{}
    ret := [][]int{}
    sort.Ints(candidates)
    for i, c := range candidates {
        if c == target {
            ret = append(ret, []int{c})
        } else if c <= target - 2 {
            cIndex[c] = i
            if c <= target - c {
                targetSets = append(targetSets, []int{c, target-c})
            }
        }
    }
    if len(cIndex) == 0 || len(targetSets) == 0 {
        return ret
    }
    for i := 0; i < len(targetSets); i++ {
        ts := targetSets[i]
        setLen := len(ts)
        if _, ok := cIndex[ts[setLen-1]]; ok {
            ret = append(ret, ts)
        }
        last := ts[setLen-2]
        for _, k := range candidates[cIndex[last]:] {
            if ts[setLen-1] - k >= k {
                newSet := make([]int, setLen+1)
                copy(newSet[:len(ts)], ts)
                newSet[setLen-1] = k
                newSet[setLen] = ts[setLen-1] - k
                targetSets = append(targetSets, newSet)
            }
        }
    }
    return ret
}

