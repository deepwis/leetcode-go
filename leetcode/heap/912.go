package leetcode

// 912. Sort an Array
//
// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity 
// and with the smallest space complexity possible.
//
// Runtime: 119 ms, faster than 96.82% of Go online submissions for Sort an Array.
// Memory Usage: 7.5 MB, less than 96.82% of Go online submissions for Sort an Array.
//
func sortArray(nums []int) []int {
    n := len(nums)
    for i := n>>1-1; i >= 0; i-- {
        maxHeap(&nums, i, n)
    }
    for i := n-1; i > 0; i-- {
        nums[0], nums[i] = nums[i], nums[0] // move max value to last position
        maxHeap(&nums, 0, i) // rebuild max heap
    }
    return nums
}

func maxHeap(nums *[]int, i, n int) {
    p := i
    for {
        l := p * 2 + 1
        if l >= n {
            break
        }
        v := l
        if l < n-1 && (*nums)[l+1] > (*nums)[v] {
            v = l+1
        }
        if (*nums)[p] >= (*nums)[v] {
            break
        }
        (*nums)[p], (*nums)[v] = (*nums)[v], (*nums)[p]
        p = v
    }
}
