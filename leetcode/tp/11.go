package leetcode

// 11. Container With Most Water
//
// You are given an integer array height of length n. There are n vertical lines drawn 
// such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container 
// contains the most water.
// Return the maximum amount of water a container can store.
//
// Runtime: 86 ms, faster than 87.43% of Go online submissions for Container With Most Water.
// Memory Usage: 8.7 MB, less than 81.87% of Go online submissions for Container With Most Water.
//
func maxArea(height []int) int {
    var max, area, w int
    left, right, low := 0, len(height) - 1, 0
    for left < right {
        w = right - left
        if height[right] < height[left] {
            low = right
            right--
            for right > left && height[right] <= height[low] {
                right--
            }
        } else {
            low = left
            left++
            for right > left && height[left] <= height[low] {
                left++
            }
        }
        area = w * height[low]
        if area > max {
            max = area
        }
    }
    return max
}

