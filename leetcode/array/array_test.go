package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 169. Majority Element
	t.Run("Test majorityElement()", func(t *testing.T) {
        inputs := [][]int{
            []int{3,2,3},
            []int{2,2,1,1,1,2,2},
            []int{6,5,5},
            []int{1},
        }
		wants := []int{3, 2, 5, 1}
        for i, in := range inputs { want := wants[i]
            got := majorityElement(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 31. Next permutation
	t.Run("Test nextPermutation()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3},
            []int{3,2,1},
            []int{1,1,5},
            []int{1,3,2},
        }
		wants := [][]int{
            []int{1,3,2},
            []int{1,2,3},
            []int{1,5,1},
            []int{2,1,3},
        }
        for i, in := range inputs {
            want := wants[i]
            nextPermutation(in)
            got := in
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 39. Combination Sum
	t.Run("Test combinationSum()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,3,6,7},
            []int{2,3,5},
            []int{2},
            []int{3,5,8},
        }
        inputs2 := []int{7,8,1,11}
		wants := [][][]int{
            [][]int{
                []int{7},
                []int{2,2,3},
            },
            [][]int{
                []int{3,5},
                []int{2,3,3},
                []int{2,2,2,2},
            },
            [][]int{},
            [][]int{
                []int{3,8},
                []int{3,3,5},
            },
        }
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := combinationSum(in, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 41. First Missing Positive
	t.Run("Test firstMissingPositive()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,0},
            []int{3,4,-1,1},
            []int{7,8,9,11,12},
            []int{0,3,8},
            []int{},
            []int{2,1,2},
        }
		wants := []int{3,2,1,1,1,3}

        for i, in := range inputs {
            want := wants[i]
            got := firstMissingPositive(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 51. N-Queens
	t.Run("Test solveNQueens()", func(t *testing.T) {
        inputs := []int{4,5}
		wants := [][][]string{
            [][]string{
                []string{".Q..","...Q","Q...","..Q."},
                []string{"..Q.","Q...","...Q",".Q.."},
            },
            [][]string{
                []string{"Q....","..Q..","....Q",".Q...","...Q."},
                []string{"Q....","...Q.",".Q...","....Q","..Q.."},
                []string{".Q...","...Q.","Q....","..Q..","....Q"},
                []string{".Q...","....Q","..Q..","Q....","...Q."},
                []string{"..Q..","Q....","...Q.",".Q...","....Q"},
                []string{"..Q..","....Q",".Q...","...Q.","Q...."},
                []string{"...Q.","Q....","..Q..","....Q",".Q..."},
                []string{"...Q.",".Q...","....Q","..Q..","Q...."},
                []string{"....Q",".Q...","...Q.","Q....","..Q.."},
                []string{"....Q","..Q..","Q....","...Q.",".Q..."},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := solveNQueens(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 84. Largest Rectangle in Histogram
	t.Run("Test largestRectangleArea()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,1,5,6,2,3},
            []int{2,4},
            []int{2,1,2},
            []int{5,4,1,2},
            []int{4,2,0,3,2,5},
        }
		wants := []int{10,4,3,8,6}

        for i, in := range inputs {
            want := wants[i]
            got := largestRectangleArea(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 88. Merge Sorted Array
	t.Run("Test merge()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,2,3,0,0,0},
                []int{2,5,6},
            },
            [][]int{
                []int{1},
                []int{},
            },
            [][]int{
                []int{0},
                []int{1},
            },
        }
        inputs2 := [][]int{
            []int{3,3},
            []int{1,0},
            []int{0,1},
        }
		wants := [][]int{
            []int{1,2,2,3,5,6},
            []int{1},
            []int{1},
        }
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            merge(in[0], in[1], in2[0], in2[1])
            got := in[0]
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 189. Rotate Array
	t.Run("Test rotate()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{1,2,3,4,5,6,7},
            []int{-1,-100,3,99},
            []int{6,5,5},
            []int{1},
            []int{1,2},
            []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27},
        }
        inputs2 := []int{3,2,3,2,3,38}
		wants := [][]int{
            []int{5,6,7,1,2,3,4},
            []int{3,99,-1,-100},
            []int{6,5,5},
            []int{1},
            []int{2,1},
            []int{17,18,19,20,21,22,23,24,25,26,27,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16},
        }
        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            rotate(in1, in2)
            got := in1
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
