package leetcode

// 236. Lowest Common Ancestor of a Binary Tree
//
// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined 
// between two nodes p and q as the lowest node in T that has both p and q as descendants 
// (where we allow a node to be a descendant of itself).”
//
// Runtime: 19 ms, faster than 65.19% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
// Memory Usage: 7.1 MB, less than 78.69% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
//
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}
