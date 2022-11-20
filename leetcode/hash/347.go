package leetcode

// 347. Top K Frequent Elements
//
// Given an integer array nums and an integer k, return the k most frequent elements. 
// You may return the answer in any order.
//
// Runtime: 16 ms, faster than 82.86% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.5 MB, less than 85.29% of Go online submissions for Top K Frequent Elements.
//
func topKFrequent(nums []int, k int) []int {
    numCnt := make(map[int]int)
    kfreq := make([]int, 0, k)
    for _, i := range nums {
        numCnt[i] += 1
    }
    buckets := make([][]int, len(nums)+1)
    for i, c := range numCnt {
        if len(buckets[c]) == 0 {
            buckets[c] = []int{i}
        } else {
            buckets[c] = append(buckets[c], i)
        }
    }
    for i := len(nums); i > 0; i-- {
        if len(buckets[i]) > 0 {
            if len(buckets[i]) + len(kfreq) > k {
                left := len(buckets[i]) - len(kfreq)
                kfreq = append(kfreq, buckets[i][:left]...)
            } else {
                kfreq = append(kfreq, buckets[i]...)
            }
            if len(kfreq) == k {
                break
            }
        }
    }
    return kfreq
}
