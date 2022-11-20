package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 4. Median of Two Sorted Arrays
	t.Run("Test findMedianSortedArrays()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{1,3},
            []int{1,2},
            []int{1,2,5},
            []int{1,2,3,4,5},
            []int{},
            []int{1,3},
        }
        inputs2 := [][]int{
            []int{2},
            []int{3,4},
            []int{3,4},
            []int{1},
            []int{1},
            []int{2,7},
        }
		wants := []float64{2.00000,2.50000, 3.00000, 2.50000, 1.00000, 2.50000}

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := findMedianSortedArrays(in1, in2)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 33. Search in Rotated Sorted Array
	t.Run("Test search()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{4,5,6,7,0,1,2},
            []int{4,5,6,7,0,1,2},
            []int{1},
            []int{8,1,2,3,4,5},
            []int{2,3,4,5,8,1},
            []int{1},
        }
        inputs2 := []int{0,3,0,2,3,1}
		wants := []int{4,-1,-1,2,1,0}

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := search(in1, in2)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 34. Find First and Last Position of Element fo Sorted Array
	t.Run("Test searchRange()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{5,7,7,8,8,10},
            []int{5,7,7,8,8,10},
            []int{},
            []int{1},
            []int{1},
        }
        inputs2 := []int{8,6,0,0,1}
		wants := [][]int{
            []int{3,4},
            []int{-1,-1},
            []int{-1,-1},
            []int{-1,-1},
            []int{0,0},
        }

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := searchRange(in1, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 35. Search Insert Position
	t.Run("Test searchInsert()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{5,7,8,10},
            []int{1,3,5,6},
            []int{1,3,5,6},
            []int{1,3,5,6},
            []int{1,3,5,6},
            []int{},
        }
        inputs2 := []int{11,5,2,7,0,5}
		wants := []int{4,2,1,4,0,0}

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := searchInsert(in1, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 74. Search a 2D Matrix
	t.Run("Test searchMatrix()", func(t *testing.T) {
        inputs1 := [][][]int{
            [][]int{
                []int{1,3,5,7},
                []int{10,11,16,20},
                []int{23,30,34,60},
            },
            [][]int{
                []int{1,3,5,7},
                []int{10,11,16,20},
                []int{23,30,34,60},
            },
            [][]int{
                []int{1,3,5,7},
                []int{10,11,16,20},
                []int{23,30,34,60},
            },
            [][]int{
                []int{1,3,5,7},
                []int{10,11,16,20},
                []int{23,30,34,60},
            },
            [][]int{
                []int{1},
            },
            [][]int{
                []int{1},
            },
        }
        inputs2 := []int{3,13,1,60,1,0}
		wants := []bool{true,false,true,true,true,false}

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := searchMatrix(in1, in2)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 153. Find Minimum in Rotated Sorted Array
	t.Run("Test findMin()", func(t *testing.T) {
        inputs := [][]int{
            []int{3,4,5,1,2},
            []int{4,5,6,7,0,1,2},
            []int{11,13,15,17},
            []int{2,1},
            []int{3,1,2},
        }
		wants := []int{1, 0, 11,1,1}
        for i, in := range inputs {
            want := wants[i]
            got := findMin(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 162. Find Peak Element
	t.Run("Test findPeakElement()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3,1},
            []int{1,2,1,3,5,6,4},
            []int{6,5},
            []int{1},
        }
		wants := [][]int{
            []int{2},
            []int{1,5},
            []int{0},
            []int{0},
        }
        for i, in := range inputs { want := wants[i]
            got := findPeakElement(in)
            if !utils.InArray(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 378. Kth Smallest Element in a Sorted Matrix
	t.Run("Test kthSmallestII()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,5,9},
                []int{10,11,13},
                []int{12,13,15},
            },
            [][]int{
                []int{-5},
            },
            [][]int{
                []int{-5,-4},
                []int{-5,-4},
            },
        }
        inputs2 := []int{8, 1, 2}
		wants := []int{13, -5, -5}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := kthSmallestII(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 704. Binary Search
	t.Run("Test searchII()", func(t *testing.T) {
        inputs := [][]int{
            []int{-1,0,3,5,9,12},
            []int{-1,0,3,5,9,12},
            []int{1},
            []int{1},
        }
        inputs2 := []int{9, 2, 1, 2}
		wants := []int{4, -1, 0, -1}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := search2(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
