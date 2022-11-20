package leetcode

import (
	"testing"
    "reflect"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 124. Binary Tree Maximum Path Sum
	t.Run("Test maxPathSum()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2}),
            utils.GenTree([]int{1,2,3}),
            utils.GenTree([]int{-10,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1}),
            utils.GenTree([]int{9,6,-3,utils.NULL,utils.NULL,-6,2,utils.NULL,utils.NULL,2,utils.NULL,-6,-6,-6}),
        }
		wants := []int{3, 6, 42, 1, 16}
        for i, in := range inputs {
            want := wants[i]
            got := maxPathSum(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 129. Sum Root to Leaf Numbers
	t.Run("Test sumNumbers()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2}),
            utils.GenTree([]int{1,2,3}),
            utils.GenTree([]int{4,9,0,5,1}),
            utils.GenTree([]int{1}),
        }
		wants := []int{12, 25, 1026, 1}
        for i, in := range inputs {
            want := wants[i]
            got := sumNumbers(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})
    // 207. Course Schedule
	t.Run("Test canFinish()", func(t *testing.T) {
        inputs1 := []int{2, 2, 3, 3, 5}
        inputs2 := [][][]int{
            [][]int{
                []int{1,0},
            },
            [][]int{
                []int{1,0},
                []int{0,1},
            },
            [][]int{
                []int{0,1},
                []int{0,2},
                []int{1,2},
            },
            [][]int{
                []int{1,0},
                []int{1,2},
                []int{0,1},
            },
            [][]int{
                []int{1,4},
                []int{2,4},
                []int{3,1},
                []int{3,2},
            },
        }
		wants := []bool{true, false, true, false, true}
        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := canFinish(in1, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 236. Lowest Common Ancestor of a Binary Tree
	t.Run("Test lowestCommonAncestor()", func(t *testing.T) {
        inputs1 := []*TreeNode{
            utils.GenTree([]int{3,5,1,6,2,0,8,utils.NULL,utils.NULL,7,4}),
            utils.GenTree([]int{3,5,1,6,2,0,8,utils.NULL,utils.NULL,7,4}),
        }
        inputs2 := [][]int{
            []int{5,4},
            []int{5,1},
        }
		wants := []int{5, 3}

        for i, in := range inputs1 {
            root := in
            p := utils.SearchNode(root, inputs2[i][0])
            q := utils.SearchNode(root, inputs2[i][1])
            want := utils.SearchNode(root, wants[i])
            got := lowestCommonAncestor(root, p, q)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 543. Diameter of Binary Tree
	t.Run("Test diameterOfBinaryTree()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2,3,4,5}),
            utils.GenTree([]int{1,2}),
        }
		wants := []int{3, 1}

        for i, in := range inputs {
            want := wants[i]
            got := diameterOfBinaryTree(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

}
