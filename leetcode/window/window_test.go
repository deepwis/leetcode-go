package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 209. Minimum Size Subarray Sum
	t.Run("Test miniSubArrayLen()", func(t *testing.T) {
        inputs := []int{7, 4, 11, 15, 213}
        inputs2 := [][]int{
            []int{2,3,1,2,4,3},
            []int{1,4,4},
            []int{1,1,1,1,1,1,1},
            []int{1,2,3,4,5},
            []int{12,28,83,4,25,26,25,2,25,25,25,12},
        }
		wants := []int{2, 1, 0, 5, 8}

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := miniSubArrayLen(in, in2)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 239. Sliding Window Maximum
	t.Run("Test maxSlidingWindow()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,3,-1,-3,5,3,6,7},
            []int{1},
            []int{1,3,1,2,0,5},
        }
        inputs2 := []int{3,1,3}
		wants := [][]int{
            []int{3,3,5,5,6,7},
            []int{1},
            []int{3,3,2,5},
        }
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := maxSlidingWindow(in, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 438. Find All Anagrams in a String
	t.Run("Test findAnagrams()", func(t *testing.T) {
        inputs := [][]string{
            []string{"cbaebabacd", "abc"},
            []string{"abab", "ab"},
        }
		wants := [][]int{
            []int{0, 6},
            []int{0, 1, 2},
        }

        for i, in := range inputs {
            want := wants[i]
            got := findAnagrams(in[0], in[1])
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })
}
