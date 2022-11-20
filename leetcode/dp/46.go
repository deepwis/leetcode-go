package leetcode

// 46. Permutations
// 
// Given an array nums of distinct integers, return all the possible permutations. 
// You can return the answer in any order.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
// Memory Usage: 2.4 MB, less than 99.62% of Go online submissions for Permutations. 
//
// Use Dynamic Programming
func permute(nums []int) [][]int {
    n := len(nums)
    all_permutations := make([][][]int, n+1)
    all_permutations[0] = [][]int{}
    all_permutations[1] = [][]int{[]int{nums[0]}}
    if n <= 1 {
        return all_permutations[n]
    }
    for i := 2; i< n+1; i++ {
        iLen := len(all_permutations[i-1]) * i
        newNum := nums[i-1]
        ilist := make([][]int, iLen)
        j := 0
        for _, v := range all_permutations[i-1] {
            for p := 0; p < i; p++ {
                newP := make([]int, 0, i)
                newP = append(newP, v[:p]...)
                newP = append(newP, newNum)
                newP = append(newP, v[p:]...)
                ilist[j] = newP
                j++
            }
        }
        all_permutations[i] = ilist
    }
    return all_permutations[n]
}
