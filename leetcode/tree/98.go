package leetcode

// 98. Validate Binary Search Tree
//
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
// * The left subtree of a node contains only nodes with keys less than the node's key.
// * The right subtree of a node contains only nodes with keys greater than the node's key.
// * Both the left and right subtrees must also be binary search trees.
//
// Runtime: 15 ms, faster than 20.57% of Go online submissions for Validate Binary Search Tree.
// Memory Usage: 5.3 MB, less than 34.22% of Go online submissions for Validate Binary Search Tree.
//
func isValidBST(root *TreeNode) bool {
    isValid, _, _ := checkBST(root)
    return isValid
}

func checkBST(root *TreeNode) (bool, int, int) {
    min, max := root.Val, root.Val
    if root.Left != nil {
        if root.Val <= root.Left.Val {
            return false, -1, -1
        }
        ok, lmin, lmax := checkBST(root.Left)
        if !ok || lmax >= root.Val {
            return false, -1, -1
        }
        min = lmin
    }
    if root.Right != nil {
        if root.Val >= root.Right.Val {
            return false, -1, -1
        }
        ok, rmin, rmax := checkBST(root.Right)
        if !ok || rmin <= root.Val {
            return false, -1, -1
        }
        max = rmax
    }
    return true, min, max
}

