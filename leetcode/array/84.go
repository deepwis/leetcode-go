package leetcode

// 84. Largest Rectangle in Histogram
//
// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, 
// return the area of the largest rectangle in the histogram.
//
// Runtime: 790 ms, faster than 6.34% of Go online submissions for Largest Rectangle in Histogram.
// Memory Usage: 8.6 MB, less than 93.66% of Go online submissions for Largest Rectangle in Histogram.
//
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    max, prevRects := heights[0], [][]int{[]int{0, heights[0]}}
    for i := 1; i < len(heights); i++ {
        if heights[i] == 0 {
            prevRects = [][]int{[]int{i,0}}
            continue
        }
        if heights[i-1] > heights[i] {
            firstLarge := len(prevRects)-1
            for j := len(prevRects)-1; j >= 0; j-- {
                if heights[i] <= prevRects[j][1] {
                    firstLarge = j
                }
            }
            prevRects[firstLarge] = []int{prevRects[firstLarge][0], heights[i]}
            prevRects = prevRects[:firstLarge+1]
        } else if heights[i-1] < heights[i] {
            prevRects = append(prevRects, []int{i, heights[i]})
        }
        for j := len(prevRects) - 1; j >= 0; j-- {
            left := prevRects[j]
            k, h := left[0], left[1]
            if h * (i-k+1) > max {
                max = h * (i-k+1)
            }
        }
    }
    return max
}

