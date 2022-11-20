package leetcode

// 295. Find Median from Data Stream
//
// The median is the middle value in an ordered integer list. If the size of the list is even, 
// there is no middle value and the median is the mean of the two middle values.
// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:
// * MedianFinder() initializes the MedianFinder object.
// * void addNum(int num) adds the integer num from the data stream to the data structure.
// * double findMedian() returns the median of all elements so far. 
//   Answers within 10-5 of the actual answer will be accepted.
//
// Runtime: 642 ms, faster than 65.93% of Go online submissions for Find Median from Data Stream.
// Memory Usage: 20.1 MB, less than 84.96% of Go online submissions for Find Median from Data Stream.
//
type MedianFinder struct {
    nums []int
}

func Constructor4() MedianFinder {
    return MedianFinder{[]int{}}
}

func (this *MedianFinder) AddNum(num int) {
    this.nums = append(this.nums, num)
    low, high := 0, len(this.nums) - 1
    for low < high {
        mid := low + (high - low) / 2
        if num > this.nums[mid] {
            low = mid+1
        } else {
            high = mid
        }
    }
    if low != len(this.nums)-1 {
        copy(this.nums[low+1:], this.nums[low:len(this.nums)-1])
        this.nums[low] = num
    }
}

func (this *MedianFinder) FindMedian() float64 {
    n := len(this.nums)
    if n % 2 == 0 {
        return (float64(this.nums[n/2-1]) + float64(this.nums[n/2])) / 2.0
    } else {
        return float64(this.nums[n/2])
    }
}

