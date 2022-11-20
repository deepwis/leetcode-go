package leetcode

// 88. Merge Sorted Array
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, 
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// Runtime: 4 ms, faster than 38.68% of Go online submissions for Merge Sorted Array.
// Memory Usage: 2.3 MB, less than 19.30% of Go online submissions for Merge Sorted Array.
// 
func merge(nums1, nums2 []int, m, n int) {
    p := m+n-1
    for i, j := m-1, n-1; j >= 0; p-- {
        if i >= 0 && nums1[i] > nums2[j] {
            nums1[i], nums1[p] = 0, nums1[i]
            i--
        } else {
            nums2[j], nums1[p] = 0, nums2[j]
            j--
        }
    }
}
