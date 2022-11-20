package leetcode

import (
    "sort"
    "strconv"
    "strings"
)

// 179. Largest Number
//
// Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
// Since the result may be very large, so you need to return a string instead of an integer.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Largest Number.
// Memory Usage: 2.3 MB, less than 98.66% of Go online submissions for Largest Number.
//
func largestNumber(nums []int) string {
    sort.Sort(LargeNumber(nums))
    if nums[0] == 0 {
        return "0"
    }
    n := 0
    for _, i := range(nums) {
        s := strconv.Itoa(i)
        n += len(s)
    }
    sb := new(strings.Builder)
    sb.Grow(n)
    for _, i := range(nums) {
        s := strconv.Itoa(i)
        sb.WriteString(s)
    }
    return sb.String()
}

type LargeNumber []int
func (this LargeNumber) Len() int { return len(this) }
func (this LargeNumber) Less(i, j int) bool {
    s1 := strconv.Itoa(this[i])
    s2 := strconv.Itoa(this[j])
    p1, p2 := &s1, &s2
    l := len(s1) + len(s2)
    k1, k2 := 0, 0
    for k := 0; k < l; k++ {
        if k == len(s1) {
            p1 = &s2
            k1 = 0
        }
        if k == len(s2) {
            p2 = &s1
            k2 = 0
        }
        if (*p1)[k1] != (*p2)[k2] {
            return (*p1)[k1] > (*p2)[k2]
        }
        k1++
        k2++
    }
    return false
}
func (this LargeNumber) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
