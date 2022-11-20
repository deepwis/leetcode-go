package leetcode

// 108. Convert Sorted Array to Binary Search Tree
//
// Given an integer array nums where the elements are sorted in ascending order, 
// convert it to a height-balanced binary search tree.
//
// Runtime: 5 ms, faster than 35.96% of Go online submissions for Convert Sorted Array to Binary Search Tree.
// Memory Usage: 3.6 MB, less than 8.94% of Go online submissions for Convert Sorted Array to Binary Search Tree.
//
func sortedArrayToBST(nums []int) *TreeNode {
    root := buildBST(&nums, 0, len(nums))
    return root
}

func buildBST(nums *[]int, l, h int) *TreeNode {
    if h - l == 1 {
        return &TreeNode{(*nums)[l], nil, nil}
    }
    mid := l + (h-l)/2
    node := &TreeNode{(*nums)[mid], nil, nil}
    node.Left = buildBST(nums, l, mid)
    if mid + 1 < h {
        node.Right = buildBST(nums, mid+1, h)
    }
    return node
}

