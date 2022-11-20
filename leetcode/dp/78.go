package leetcode

// 78. Subsets
//
// Given an integer array nums of unique elements, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Runtime: 4 ms, faster than 23.23% of Go online submissions for Subsets.
// Memory Usage: 2.3 MB, less than 80.65% of Go online submissions for Subsets.
//
func subsets(nums []int) [][]int {
    ss := [][]int{[]int{}}
    if len(nums) == 0 {
        return ss
    }
    for _, v := range nums {
        newSets := make([][]int, len(ss))
        for i, vv := range ss {
            l := len(vv)+1
            newSets[i] = make([]int, l)
            // use newSets[i] = append(vv, v) or newSets[i] = append(ss[i],v) will cause a bug
            // nums = [9,0,3,5,7], ss[i] = [9,0,3,5], v = 7, this will get a wrong result [9, 0, 3, 7, 7]
            copy(newSets[i][:l-1], vv)
            newSets[i][l-1] = v
        }
        ss = append(ss, newSets...)
        newSets = [][]int{}
    }
    return ss
}

// Runtime: 4 ms, faster than 23.23% of Go online submissions for Subsets.
// Memory Usage: 2.3 MB, less than 80.65% of Go online submissions for Subsets.
//
func subsets2(nums []int) [][]int {
    ss := [][]int{[]int{}}
    if len(nums) == 0 {
        return ss
    }
    for _, v := range nums {
        l := len(ss)
        for i := 0; i < l; i++ {
            newSet := make([]int, len(ss[i])+1)
            copy(newSet[:len(ss[i])], ss[i])
            newSet[len(ss[i])] = v
            ss = append(ss, newSet)
        }
    }
    return ss
}
