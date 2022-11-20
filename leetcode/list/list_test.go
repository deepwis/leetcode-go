package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 2. Add Two Numbers
	t.Run("Test addTwoNumbers()", func(t *testing.T) {
        inputs := [][]*ListNode{
            []*ListNode{
                utils.GenSingly([]int{2,4,3}),
                utils.GenSingly([]int{5,6,4}),
            },
            []*ListNode{
                utils.GenSingly([]int{0}),
                utils.GenSingly([]int{0}),
            },
            []*ListNode{
                utils.GenSingly([]int{9,9,9,9,9,9,9}),
                utils.GenSingly([]int{9,9,9,9}),
            },
            []*ListNode{
                utils.GenSingly([]int{9,9,1}),
                utils.GenSingly([]int{1}),
            },
        }
		wants := [][]int{
            []int{7,0,8},
            []int{0},
            []int{8,9,9,9,0,0,0,1},
            []int{0,0,2},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(addTwoNumbers(in[0], in[1]))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 206. Reverse Linked List
	t.Run("Test reverseList()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{1,2}),
            utils.GenSingly([]int{1}),
        }
		wants := [][]int{
            []int{5,4,3,2,1},
            []int{2,1},
            []int{1},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(reverseList(in))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
            }
        }
	})

    // 21. Merge Two Sorted Lists
	t.Run("Test mergeTwoLists()", func(t *testing.T) {
        inputs := [][]*ListNode{
            []*ListNode{
                utils.GenSingly([]int{1,2,4}),
                utils.GenSingly([]int{1,3,4}),
            },
            []*ListNode{
                utils.GenSingly([]int{}),
                utils.GenSingly([]int{0}),
            },
            []*ListNode{
                utils.GenSingly([]int{-9,3}),
                utils.GenSingly([]int{5,7}),
            },
            []*ListNode{
                utils.GenSingly([]int{5,9,11}),
                utils.GenSingly([]int{-1,1}),
            },
        }
		wants := [][]int{
            []int{1,1,2,3,4,4},
            []int{0},
            []int{-9,3,5,7},
            []int{-1,1,5,9,11},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(mergeTwoLists(in[0], in[1]))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 234. Palindrome Linked List
	t.Run("Test isPalindrome()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,2,1}),
            utils.GenSingly([]int{1,2}),
        }
		wants := []bool{true, false}

        for i, in := range inputs {
            want := wants[i]
            got := isPalindrome(in)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 24. Swap Nodes in Pairs
	t.Run("Test swapPairs()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,}),
            utils.GenSingly([]int{1}),
            utils.GenSingly([]int{1,2,3,4,5}),
        }
		wants := [][]int{
            []int{2,1,4,3},
            []int{1},
            []int{2,1,4,3,5},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(swapPairs(in))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 25. Reverse Nodes in k-Group
	t.Run("Test reverseKGroup()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,}),
            utils.GenSingly([]int{1}),
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{1,2,3,4,5}),
        }
        inputs2 := []int{2,2,2,3}
		wants := [][]int{
            []int{2,1,4,3},
            []int{1},
            []int{2,1,4,3,5},
            []int{3,2,1,4,5},
        }

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := utils.ToArray(reverseKGroup(in, in2))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 83. Remove Duplicates from Sorted List
	t.Run("Test deleteDuplicates()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,1,2,3,3}),
            utils.GenSingly([]int{1}),
            utils.GenSingly([]int{1,1,2}),
        }
		wants := [][]int{
            []int{1,2,3},
            []int{1},
            []int{1,2},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(deleteDuplicates(in))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

}
