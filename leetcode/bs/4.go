package leetcode

// 4. Median of Two Sorted Arrays
//
// Given two sorted arrays nums1 and nums2 of size m and n respectively, 
// return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
//
// Runtime: 18 ms, faster than 82.18% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 5 MB, less than 97.04% of Go online submissions for Median of Two Sorted Arrays.
//
func findMedianSortedArrays(nums1, nums2 []int) float64 {
    var mid, next int
    m, n := len(nums1), len(nums2)
    k, sorted := (m + n + 1) / 2, 0
    low1, low2 := 0, 0

    if m == 0 && n == 0 {
        return 0.0
    }
    if m == 0 {
        if n % 2 == 0 {
            return float64(nums2[k-1] + nums2[k]) / 2
        } else {
            return float64(nums2[k-1])
        }
    }
    if n == 0 {
        if m % 2 == 0 {
            return float64(nums1[k-1] + nums1[k]) / 2
        } else {
            return float64(nums1[k-1])
        }
    }
    for i := k / 2; i > 0; i = (k - sorted) / 2 {
        high1, high2 := low1 + i, low2 + i
        if high1 >= m {
            high1 = m
        }
        if high2 >= n {
            high2 = n
        }

        if nums1[high1-1] <= nums2[high2-1] {
            sorted += high1 - low1
            low1 = high1
        } else {
            sorted += high2 - low2
            low2 = high2
        }
        if low1 == m || low2 == n {
            break
        }
    }

    if low1 >= m {
        low2 = low2 + (k - sorted - 1)
        mid, next = nums2[low2], nums2[low2+1]
    } else if low2 >= n {
        low1 = low1 + (k - sorted - 1)
        mid, next = nums1[low1], nums1[low1+1]
    } else {
        if nums1[low1] <= nums2[low2] {
            mid = nums1[low1]
            if low1 + 1 < m && nums1[low1+1] <= nums2[low2] {
                next = nums1[low1+1]
            } else {
                next = nums2[low2]
            }
        } else {
            mid = nums2[low2]
            if low2 + 1 < n && nums1[low1] > nums2[low2+1] {
                next = nums2[low2+1]
            } else {
                next = nums1[low1]
            }
        }
    }

    if (m + n) % 2 == 0 {
        return float64(mid + next) / 2
    } else {
        return float64(mid)
    }
}

