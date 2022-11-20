package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 23. Merge K Sorted Lists
	t.Run("Test mergeKLists()", func(t *testing.T) {
        inputs := [][]*ListNode{
            []*ListNode{
                utils.GenSingly([]int{1,4,5}),
                utils.GenSingly([]int{1,3,4}),
                utils.GenSingly([]int{2,6}),
            },
            []*ListNode{
                utils.GenSingly([]int{1,4,5}),
                utils.GenSingly([]int{}),
                utils.GenSingly([]int{2,3}),
                utils.GenSingly([]int{6,7}),
            },
        }
		wants := [][]int{
            []int{1,1,2,3,4,4,5,6},
            []int{1,2,3,4,5,6,7},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(mergeKLists(in))
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
}
