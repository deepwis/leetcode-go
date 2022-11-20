package leetcode

import "container/heap"

// 23. Merge K Sorted Lists
//
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Runtime: 9 ms, faster than 90.73% of Go online submissions for Merge k Sorted Lists.
// Memory Usage: 5.3 MB, less than 95.09% of Go online submissions for Merge k Sorted Lists.
//
func mergeKLists(lists []*ListNode) *ListNode {
    var merged *ListNode
    var pq PriorityQueue

    i0 := 0
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {
            continue
        }
        if i0 < i {
            lists[i0], lists[i] = lists[i], nil
        }
        i0++
    }
    if i0 == 0 {
        return nil
    }
    if i0 == 1 {
        return lists[0]
    }

    lists = lists[0 : i0]

    pq = lists

    heap.Init(&pq)
    merged = heap.Pop(&pq).(*ListNode)
    cur := merged
    for pq.Len() > 0 {
        x := heap.Pop(&pq).(*ListNode)
        for cur.Next != nil && cur.Next.Val <= x.Val {
            cur = cur.Next
        }
        if cur.Next != nil {
            heap.Push(&pq, cur.Next)
        }
        cur.Next = x
        cur = cur.Next
    }
    return merged
}

type PriorityQueue []*ListNode

func (q PriorityQueue) Len() int {
    return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
    return q[i].Val < q[j].Val
}

func (q PriorityQueue) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
    *q = append(*q, x.(*ListNode))
}

func (q *PriorityQueue) Pop() interface{} {
    old := *q
    n := len(old)
    x := old[n-1]
    *q = old[: n-1]
    return x
}

