package leetcode

// 42. Trapping Rain Water
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, 
// compute how much water it can trap after raining.
//
// Runtime: 21 ms, faster than 61.96% of Go online submissions for Trapping Rain Water.
// Memory Usage: 5.4 MB, less than 74.08% of Go online submissions for Trapping Rain Water.
//
func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    l, r := 0, len(height) - 1
    for l < r - 1 {
        if height[l] == 0 || height[l] <= height[l+1] {
            l++
        } else if height[r] == 0 || height[r] <= height[r-1] {
            r--
        } else {
            break
        }
    }

    if l >= r - 1 {
        return 0
    }

    area, filled := 0, 0
    for l < r {
        if height[l] <= height[r] {
            for i := l + 1; i <= r; i++ {
                if height[i] < height[l] {
                    filled += height[i]
                } else {
                    area += height[l] * (i - l - 1) - filled
                    filled = 0
                    l = i
                }
                if height[i] > height[r] {
                    break
                }
            }
        }
        if height[l] > height[r] {
            for i := r - 1; i >= l; i-- {
                if height[i] < height[r] {
                    filled += height[i]
                } else {
                    area += height[r] * (r - i - 1) - filled
                    filled = 0
                    r = i
                }
                if height[i] > height[l] {
                    break
                }
            }
        }
    }
    return area
}

