package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 11. Container With Most Water
	t.Run("Test maxArea()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,8,6,2,5,4,8,3,7},
            []int{4,4,2,11,0,11,5,11,13,8},
            []int{1,1},
            []int{8},
            []int{},
        }
		wants := []int{49,55,1,0,0}

        for i, in := range inputs {
            want := wants[i]
            got := maxArea(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 19. Remove Nth Node from End of List
	t.Run("Test removeNthFromEnd()", func(t *testing.T) {
        inputs1 := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{1,2}),
            utils.GenSingly([]int{1,2}),
        }
        inputs2 := []int{2,2,1}
		wants := [][]int{
            []int{1,2,3,5},
            []int{2},
            []int{1},
        }

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := utils.ToArray(removeNthFromEnd(in1, in2))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 92. Reverse List II
	t.Run("Test reverseBetween()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{5}),
            utils.GenSingly([]int{1,2,3,4,5}),
        }
        inputs2 := [][]int{
            []int{2,4},
            []int{1,1},
            []int{1,5},
        }
		wants := [][]int{
            []int{1,4,3,2,5},
            []int{5},
            []int{5,4,3,2,1},
        }

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := utils.ToArray(reverseBetween(in, in2[0], in2[1]))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 121. Best Time to Buy and Sell Stock
	t.Run("Test maxProfit()", func(t *testing.T) {
        inputs := [][]int{
            []int{7,1,5,3,6,4},
            []int{7,6,4,3,1},
            []int{1,2,3,4,5},
        }
		wants := []int{5,0,4}

        for i, in := range inputs {
            want := wants[i]
            got := maxProfit2(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
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

    // 238. Product of Array Except Self
	t.Run("Test productExceptSelf()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3,4},
            []int{-1,1,0,-3,3},
            []int{1,5},
        }
		wants := [][]int{
            []int{24,12,8,6},
            []int{0,0,9,0,0},
            []int{5,1},
        }
        for i, in := range inputs {
            want := wants[i]
            got := productExceptSelf(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 283. Move Zeroes
	t.Run("Test moveZeroes()", func(t *testing.T) {
        inputs := [][]int{
            []int{0,1,0,3,12},
            []int{0},
            []int{0,0},
        }
		wants := [][]int{
            []int{1,3,12,0,0},
            []int{0},
            []int{0,0},
        }

        for i, in := range inputs {
            want := wants[i]
            moveZeroes(in)
            got := in
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
            }
        }
    })

    // 42. Trapping Rain Water
	t.Run("Test trap()", func(t *testing.T) {
        inputs := [][]int{
            []int{0,1,0,2,1,0,1,3,2,1,2,1},
            []int{4,2,0,3,2,5},
            []int{1,2},
            []int{0,3,8},
            []int{},
            []int{2,0,2},
        }
		wants := []int{6,9,0,0,0,2}

        for i, in := range inputs {
            want := wants[i]
            got := trap(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 53. Maximum Subarray
	t.Run("Test maxSubArray()", func(t *testing.T) {
        inputs := [][]int{
            []int{-2,1,-3,4,-1,2,1,-5,4},
            []int{1},
            []int{5,4,-1,7,8},
            []int{-1, -3, -5},
            []int{0, -1, -3},
        }
		wants := []int{6,1,23,-1,0}

        for i, in := range inputs {
            want := wants[i]
            got := maxSubArray(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 75. Sort Colors
	t.Run("Test sortColors()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,0,2,1,1,0},
            []int{2,0,1},
            []int{1,1,0},
            []int{2,0,2},
            []int{2},
            []int{1,2,2,2,2,0,0,0,1,1},
            []int{1,1,0,0,0,2,2,2,1,1},
        }
		wants := [][]int{
            []int{0,0,1,1,2,2},
            []int{0,1,2},
            []int{0,1,1},
            []int{0,2,2},
            []int{2},
            []int{0,0,0,1,1,1,2,2,2,2},
            []int{0,0,0,1,1,1,1,2,2,2},
        }

        for i, in := range inputs {
            want := wants[i]
            sortColors(in)
            got := in
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 82. Remove Dumplicates from Sorted List II
	t.Run("Test deleteDuplicates()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,3,4,4,5}),
            utils.GenSingly([]int{1}),
            utils.GenSingly([]int{1,1,1,2,3}),
            utils.GenSingly([]int{1,2,2}),
        }
		wants := [][]int{
            []int{1,2,5},
            []int{1},
            []int{2,3},
            []int{1},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(deleteDuplicates(in))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 103. Binary Tree Zigzag Level Order Traversal
	t.Run("Test zigzagLevelOrder()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1}),
            utils.GenTree([]int{1,2,3,4,utils.NULL,6,7,utils.NULL,utils.NULL,utils.NULL,utils.NULL,8,utils.NULL,9}),
        }
		wants := [][][]int{
            [][]int{
                []int{3},
                []int{20,9},
                []int{15,7},
            },
            [][]int{
                []int{1},
            },
            [][]int{
                []int{1},
                []int{3,2},
                []int{4,6,7},
                []int{8},
                []int{9},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := zigzagLevelOrder(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 143. Reorder List
	t.Run("Test reorderList()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4}),
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{1,2,3,4,5,6}),
        }
		wants := [][]int{
            []int{1,4,2,3},
            []int{1,5,2,4,3},
            []int{1,6,2,5,3,4},
        }

        for i, in := range inputs {
            want := wants[i]
            reorderList(in)
            got := utils.ToArray(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 763. Partition Labels
	t.Run("Test partitionLabels()", func(t *testing.T) {
        inputs := []string{
            "ababcbacadefegdehijhklij",
            "eccbbbbdec",
            "aa",
            "vhaagbqkaq",
        }
		wants := [][]int{
            []int{9,7,8},
            []int{10},
            []int{2},
            []int{1,1,8},
        }

        for i, in := range inputs {
            want := wants[i]
            got := partitionLabels(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 876. Middle of the Linked List
	t.Run("Test middleNode()", func(t *testing.T) {
        inputs := []*ListNode{
            utils.GenSingly([]int{1,2,3,4,5}),
            utils.GenSingly([]int{1,2,3,4,5,6}),
        }
		wants := [][]int{
            []int{3,4,5},
            []int{4,5,6},
        }

        for i, in := range inputs {
            want := wants[i]
            got := utils.ToArray(middleNode(in))
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 160. Intersection of Two Linked Lists
}
