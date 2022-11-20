package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 1337. The K Weakest Rows in a Matrix
	t.Run("Test kWeakestRows()", func(t *testing.T) {
        input := [][]int{
            []int{1,1,0,0,0},
            []int{1,1,1,1,0},
            []int{1,0,0,0,0},
            []int{1,1,0,0,0},
            []int{1,1,1,1,1},
        }
		want := []int{2, 0, 3}
        got := kWeakestRows(input, 3)
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }

        input = [][]int{
            []int{1,0,0,0},
            []int{1,1,1,1},
            []int{1,0,0,0},
            []int{1,0,0,0},
        }
        want = []int{0,2}
        got = kWeakestRows(input, 2)
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }

        input = [][]int{
            []int{1,0},
            []int{1,0},
            []int{1,0},
            []int{1,1},
        }
        want = []int{0,1,2,3}
        got = kWeakestRows(input, 4)
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
	})

    // 240. Search a 2D Matrix II
	t.Run("Test searchMatrixII()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,4,7,11,15},
                []int{2,5,8,12,19},
                []int{3,6,9,16,22},
                []int{10,13,14,17,24},
                []int{18,21,23,26,30},
            },
            [][]int{
                []int{1,4,7,11,15},
                []int{2,5,8,12,19},
                []int{3,6,9,16,22},
                []int{10,13,14,17,24},
                []int{18,21,23,26,30},
            },
            [][]int{
                []int{1},
            },
            [][]int{
                []int{1},
            },
        }
        inputs2 := []int{5, 20, 1, 0}
		wants := []bool{true, false, true, false}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := searchMatrixII(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 48. Rotate Image
	t.Run("Test rotate()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
            [][]int{
                []int{5,1,9,11},
                []int{2,4,8,10},
                []int{13,3,6,7},
                []int{15,14,12,16},
            },
            [][]int{
                []int{1,2},
                []int{3,4},
            },
            [][]int{
                []int{3},
            },
        }
		wants := [][][]int{
            [][]int{
                []int{7,4,1},
                []int{8,5,2},
                []int{9,6,3},
            },
            [][]int{
                []int{15,13,2,5},
                []int{14,3,4,1},
                []int{12,6,8,9},
                []int{16,7,10,11},
            },
            [][]int{
                []int{3,1},
                []int{4,2},
            },
            [][]int{
                []int{3},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            rotate(in)
            got := in
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 54. Spiral Matrix
	t.Run("Test spiralOrder()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
            [][]int{
                []int{5,1,9,11},
                []int{2,4,8,10},
                []int{13,3,6,7},
                []int{15,14,12,16},
            },
            [][]int{
                []int{1,2},
                []int{3,4},
            },
            [][]int{
                []int{3},
            },
            [][]int{
                []int{1,2,3,4},
                []int{5,6,7,8},
                []int{9,10,11,12},
            },
        }
		wants := [][]int{
            []int{1,2,3,6,9,8,7,4,5},
            []int{5,1,9,11,10,7,16,12,14,15,13,2,4,8,6,3},
            []int{1,2,4,3},
            []int{3},
            []int{1,2,3,4,8,12,11,10,9,5,6,7},
        }

        for i, in := range inputs {
            want := wants[i]
            got := spiralOrder2(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 498. Diagonal Traverse
	t.Run("Test findDiagonalOrder()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
            [][]int{
                []int{1,2},
                []int{3,4},
            },
            [][]int{
                []int{3},
            },
        }
		wants := [][]int{
            []int{1,2,4,7,5,3,6,8,9},
            []int{1,2,3,4},
            []int{3},
        }

        for i, in := range inputs {
            want := wants[i]
            got := findDiagonalOrder(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

}
