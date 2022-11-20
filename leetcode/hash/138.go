package leetcode

// 138. Copy List with Random Pointer
//
// A linked list of length n is given such that each node contains an additional random pointer, 
// which could point to any node in the list, or null.
// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, 
// where each new node has its value set to the value of its corresponding original node. 
// Both the next and random pointer of the new nodes should point to new nodes in the copied list 
// such that the pointers in the original list and copied list represent the same list state. 
// None of the pointers in the new list should point to nodes in the original list.
// For example, if there are two nodes X and Y in the original list, where X.random --> Y, 
// then for the corresponding two nodes x and y in the copied list, x.random --> y.
// Return the head of the copied linked list.
// The linked list is represented in the input/output as a list of n nodes. 
// Each node is represented as a pair of [val, random_index] where:
// * val: an integer representing Node.val
// * random_index: the index of the node (range from 0 to n-1) that the random pointer points to, 
// * or null if it does not point to any node.
// Your code will only be given the head of the original linked list.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Copy List with Random Pointer.
// Memory Usage: 3.7 MB, less than 13.85% of Go online submissions for Copy List with Random Pointer.
//
type Node struct {
    Val int
    Next *Node
    Random *Node
}
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    sourceToTarget := make(map[*Node]*Node)
    randomToTargets := make(map[*Node][]*Node)
    newHead := &Node{head.Val, nil, nil}
    cur := head
    target := newHead
    for cur != nil {
        sourceToTarget[cur] = target
        if targets, ok := randomToTargets[cur]; ok {
            for _, t := range targets {
                t.Random = target
            }
        }
        if cur.Random != nil {
            if v, ok := sourceToTarget[cur.Random]; ok {
                target.Random = v
            } else if _, ok := randomToTargets[cur.Random]; ok {
                randomToTargets[cur.Random] = append(randomToTargets[cur.Random], target)
            } else {
                randomToTargets[cur.Random] = []*Node{target}
            }
        }
        if cur.Next != nil {
            target.Next = &Node{cur.Next.Val, nil, nil}
        }
        cur = cur.Next
        target = target.Next
    }
    return newHead
}
