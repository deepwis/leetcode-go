package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 15. 3Sum
	t.Run("Test threeSum()", func(t *testing.T) {
        inputs := [][]int{
            []int{-1,0,1,2,-1,-4},
            []int{0,0,0,0,0},
            []int{2,6,3,4,1,-8,5,0,-3},
        }
		wants := [][][]int{
            [][]int{[]int{-1, -1, 2}, []int{-1,0,1}},
            [][]int{[]int{0,0,0}},
            [][]int{[]int{-8,2,6},[]int{-8,3,5},[]int{-3,0,3},[]int{-3,1,2}},
        }

        for i, in := range inputs {
            want := wants[i]
            got := threeSum(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 179. Largest Number 
	t.Run("Test largestNumber()", func(t *testing.T) {
        inputs := [][]int{
            []int{10,2},
            []int{3,30,34,5,9},
            []int{6,5,5},
            []int{1},
            []int{432,43243},
        }
		wants := []string{"210", "9534330", "655", "1", "43243432"}
        for i, in := range inputs {
            want := wants[i]
            got := largestNumber(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 215. Kth Largest Element in an Array
	t.Run("Test findKthLargest()", func(t *testing.T) {
        inputs := [][]int{
            []int{3,2,1,5,6,4},
            []int{3,2,3,1,2,4,5,5,6},
            []int{6,5,5},
            []int{1},
        }
        inputs2 := []int{2, 4, 3, 1}
		wants := []int{5, 4, 5, 1}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := findKthLargest(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 56. Merge Intervals
	t.Run("Test merge()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,3},
                []int{15,18},
                []int{2,6},
                []int{8,10},
            },
            [][]int{
                []int{1,4},
                []int{4,5},
            },
            [][]int{
                []int{1,4},
                []int{5,6},
            },
        }
		wants := [][][]int{
            [][]int{
                []int{1,6},
                []int{8,10},
                []int{15,18},
            },
            [][]int{
                []int{1,5},
            },
            [][]int{
                []int{1,4},
                []int{5,6},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := merge(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 912. Sort an Array
	t.Run("Test sortArray()", func(t *testing.T) {
        inputs := [][]int{
            []int{5,2,3,1},
            []int{5,1,1,2,0,0},
            []int{2,6},
        }
		wants := [][]int{
            []int{1,2,3,5},
            []int{0,0,1,1,2,5},
            []int{2,6},
        }

        for i, in := range inputs {
            want := wants[i]
            got := sortArray(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 148. Sort List
}
