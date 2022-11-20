package leetcode

import (
    "strings"
    "strconv"
)

// 297. Serialize and Deserialize Binary Tree
//
// Serialization is the process of converting a data structure or object into a sequence of bits 
// so that it can be stored in a file or memory buffer, or transmitted across a network connection link 
// to be reconstructed later in the same or another computer environment.
// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on 
// how your serialization/deserialization algorithm should work. You just need to ensure that 
// a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.
// Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not 
// necessarily need to follow this format, so please be creative and come up with different approaches yourself.
//
// Runtime: 13 ms, faster than 90.50% of Go online submissions for Serialize and Deserialize Binary Tree.
// Memory Usage: 7.2 MB, less than 72.73% of Go online submissions for Serialize and Deserialize Binary Tree.
//
type Codec struct {
    sep byte
}

func Constructor() Codec {
    return Codec{sep: ','}
}

func (this *Codec) Serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    sb := new(strings.Builder)
    q := []*TreeNode{root}
    for len(q) > 0 {
        n := len(q)
        for i := 0; i < n; i++ {
            if q[i] != nil {
                s := strconv.Itoa(q[i].Val)
                sb.WriteString(s)
                sb.Grow(len(s))
                q = append(q, q[i].Left)
                q = append(q, q[i].Right)
            }
            sb.Grow(1)
            sb.WriteByte(this.sep)
        }
        q = q[n:]
    }
    s := sb.String()
    k := len(s)-1
    for k >= 0 && s[k] == this.sep {
        k--
    }
    return s[:k+1]
}

func (this *Codec) Deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    nodes := make([]*TreeNode, 1)
    low, high := 0, 0
    for high <= len(data) {
        if high == len(data) || data[high] == ',' {
            if low < high {
                v, _ := strconv.Atoi(data[low:high])
                nodes[0] = &TreeNode{Val: v} // root
                high++
                low = high
                break
            }
        }
        high++
    }
    root := nodes[0]
    var p *TreeNode
    var n *TreeNode
    for high < len(data) {
        if data[high] != ',' {
            high++
            continue
        }
        if low < high {
            v, _ := strconv.Atoi(data[low:high])
            n = &TreeNode{Val: v}
            nodes = append(nodes, n)
        }
        if p == nil {
            p = nodes[0]
            nodes = nodes[1:]
            if n != nil {
                p.Left = n
                n = nil
            }
        } else {
            if n != nil {
                p.Right = n
                n = nil
            }
            p = nil
        }
        high++
        low = high
    }
    if high > low {
        v, _ := strconv.Atoi(data[low:high])
        n = &TreeNode{Val: v}
        if p == nil {
            p = nodes[0]
            p.Left = n
        } else {
            p.Right = n
        }
    }
    return root
}
