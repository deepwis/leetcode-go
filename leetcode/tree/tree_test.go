package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 101. Symmetric Tree
	t.Run("Test isSymmetric()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2,2,3,4,4,3}),
            utils.GenTree([]int{1,2,2,utils.NULL,3,utils.NULL,3}),
        }
		wants := []bool{true, false}

        for i, in := range inputs {
            want := wants[i]
            got := isSymmetric(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 102. Binary Tree Level Order Traversal
	t.Run("Test levelOrder()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1}),
            utils.GenTree([]int{1,2,3,4,utils.NULL,6,7,utils.NULL,utils.NULL,utils.NULL,utils.NULL,8,utils.NULL,9}),
        }
		wants := [][][]int{
            [][]int{
                []int{3},
                []int{9,20},
                []int{15,7},
            },
            [][]int{
                []int{1},
            },
            [][]int{
                []int{1},
                []int{2,3},
                []int{4,6,7},
                []int{8},
                []int{9},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := levelOrder2(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 104. Maximum Depth of Binary Tree
	t.Run("Test maxDepth()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1,utils.NULL,2}),
            utils.GenTree([]int{1}),
            utils.GenTree([]int{}),
        }
		wants := []int{3,2,1,0}

        for i, in := range inputs {
            want := wants[i]
            got := maxDepth2(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 105. Construct Binary Tree from Preorder and Inorder Traversal
	t.Run("Test buildTree()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{3,9,20,15,7},
                []int{9,3,15,20,7},
            },
            [][]int{
                []int{-1},
                []int{-1},
            },
        }
		wants := [][]int{
            []int{3,9,20,utils.NULL,utils.NULL,15,7},
            []int{-1},
        }

        for i, in := range inputs {
            want := wants[i]
            got := []int{}
            utils.TreeToArray(buildTree(in[0], in[1]), &got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 110. Balanced Binary Tree
	t.Run("Test isBalanced()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1,2,2,3,3,utils.NULL,utils.NULL,4,4}),
        }
		wants := []bool{true,false}

        for i, in := range inputs {
            want := wants[i]
            got := isBalanced(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 112. Path Sum
	t.Run("Test hasPathSum()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{5,4,8,11,utils.NULL,13,4,7,2,utils.NULL,utils.NULL,5,1}),
            utils.GenTree([]int{1,2,3}),
            utils.GenTree([]int{1,2}),
        }
        inputs2 := []int{22, 5, 0}
		wants := []bool{true, false, false}

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := hasPathSum(in, in2)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 113. Path Sum II
	t.Run("Test pathSum()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{5,4,8,11,utils.NULL,13,4,7,2,utils.NULL,utils.NULL,5,1}),
            utils.GenTree([]int{1,2,3}),
            utils.GenTree([]int{1,2}),
        }
        inputs2 := []int{22, 5, 0}
		wants := [][][]int{
            [][]int{
                []int{5,4,11,2},
                []int{5,8,4,5},
            },
            [][]int{},
            [][]int{},
        }

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := pathSum(in, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 114. Flatten Binary Tree to Linked List
	t.Run("Test flatten()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2,5,3,4,utils.NULL,6}),
            utils.GenTree([]int{0}),
        }
		wants := [][]int{
            []int{1,utils.NULL,2,utils.NULL,3,utils.NULL,4,utils.NULL,5,utils.NULL,6},
            []int{0},
        }

        for i, in := range inputs {
            want := wants[i]
            flatten(in)
            got := []int{}
            utils.TreeToArray(in, &got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

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

    // 199. Binary Tree Right Side View
	t.Run("Test rightSideView()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2,3,utils.NULL,5,utils.NULL,4}),
            utils.GenTree([]int{1,utils.NULL,3}),
            utils.GenTree([]int{}),
        }
		wants := [][]int{
            []int{1,3,4},
            []int{1,3},
            []int{},
        }
        for i, in := range inputs {
            want := wants[i]
            got := rightSideView(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 226. Invert Binary Tree
	t.Run("Test invertTree()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{4,2,7,1,3,6,9}),
            utils.GenTree([]int{2,1,3}),
            utils.GenTree([]int{}),
        }
		wants := [][]int{
            []int{4,7,2,9,6,3,1},
            []int{2,3,1},
            []int{},
        }

        for i, in := range inputs {
            want := wants[i]
            got :=[]int{}
            utils.TreeToArray(invertTree(in), &got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 230. Kth Smallest Element in BST
	t.Run("Test kthSmallest()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,1,4,utils.NULL,2}),
            utils.GenTree([]int{5,3,6,2,4,utils.NULL,utils.NULL,1}),
            utils.GenTree([]int{1}),
        }
        inputs2 := []int{1, 3, 1}
		wants := []int{1, 3, 1}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := kthSmallest(in, in2)
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

    // 437. Path Sum III
	t.Run("Test pathSumIII()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{10,5,-3,3,2,utils.NULL,11,3,-2,utils.NULL,1}),
            utils.GenTree([]int{5,4,8,11,utils.NULL,13,4,7,2,utils.NULL,utils.NULL,5,1}),
        }
        inputs2 := []int{8,22}
		wants := []int{3, 3}

        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := pathSumIII(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 94. Binary Tree Inorder Traversal
	t.Run("Test inorderTraversal()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,0,2,0,0,3}),
            utils.GenTree([]int{1}),
        }
		wants := [][]int{
            []int{0,0,0,1,3,2},
            []int{1},
        }

        for i, in := range inputs {
            want := wants[i]
            got := inorderTraversal(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 98. Validate Binary Search Tree
	t.Run("Test isValidBST()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{2,1,3}),
            utils.GenTree([]int{5,1,4,utils.NULL,utils.NULL,3,6}),
            utils.GenTree([]int{2}),
            utils.GenTree([]int{2,2,2}),
            utils.GenTree([]int{5,4,6,utils.NULL,utils.NULL,3,7}),
        }
		wants := []bool{true,false,true,false,false}

        for i, in := range inputs {
            want := wants[i]
            got := isValidBST(in)
            if got != want {
                arr := []int{}
                utils.TreeToArray(in, &arr)
                t.Errorf("got: %v (%T), want: %v (%T), input: %v(%T)", got, got, want, want, arr, in)
            }
        }
    })

    // 144. Binary Tree Preorder Traversal
	t.Run("Test preorderTraversal()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,utils.NULL,2,3}),
            utils.GenTree([]int{1}),
        }
		wants := [][]int{
            []int{1,2,3},
            []int{1},
        }

        for i, in := range inputs {
            want := wants[i]
            got := preorderTraversal(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 662. Maximum Width of Binary Tree
	t.Run("Test widthOfBinaryTree()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,3,2,5,3,utils.NULL,9}),
            utils.GenTree([]int{1,3,2,5,utils.NULL,utils.NULL,9,6,utils.NULL,7}),
            utils.GenTree([]int{1,3,2,5}),
        }
		wants := []int{4,7,2}

        for i, in := range inputs {
            want := wants[i]
            got := widthOfBinaryTree(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 958. Chcek Completeness of a Binary Tree
	t.Run("Test isCompleteTree()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{1,2,3,4,5,6}),
            utils.GenTree([]int{1,2,3,4,5,utils.NULL,7}),
        }
		wants := []bool{true, false}

        for i, in := range inputs {
            want := wants[i]
            got := isCompleteTree(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 108. Convert Sorted Array to Binary Search Tree
}
