package leetcode

// 236. Lowest Common Ancestor of a Binary Tree
//
// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined 
// between two nodes p and q as the lowest node in T that has both p and q as descendants 
// (where we allow a node to be a descendant of itself).”
//
// Runtime: 41 ms, faster than 7.72% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
// Memory Usage: 7.8 MB, less than 33.86% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
//
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pParents, qParents := []*TreeNode{}, []*TreeNode{}
    ret := root
    nodePath(root, p, &pParents)
    nodePath(root, q, &qParents)
    pLen, qLen := len(pParents), len(qParents)
    for i := 2; i <= pLen && i <= qLen; i++ {
        if pParents[pLen-i] != qParents[qLen-i] {
            break
        }
        ret = pParents[pLen-i]
    }
    return ret
}

func nodePath(root, n *TreeNode, parents *[]*TreeNode) bool {
    if root == n {
        *parents = append(*parents, root)
        return true
    }
    if root.Left != nil && nodePath(root.Left, n, parents) {
        *parents = append(*parents, root)
        return true
    }
    if root.Right != nil && nodePath(root.Right, n, parents) {
        *parents = append(*parents, root)
        return true
    }
    return false
}
