package leetcode

import (
	"testing"
)

// 295. Find Median from Data Stream
func TestMedianFinder(t *testing.T) {
	t.Run("Test all()", func(t *testing.T) {
        f := Constructor4()
        f.AddNum(-1)
        f.AddNum(-2)
        f.AddNum(-3)
        want := -2.0
        got := f.FindMedian()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
}

