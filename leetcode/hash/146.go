package leetcode

// 146. LRU Cache
//
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
// * LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// * int get(int key) Return the value of the key if the key exists, otherwise return -1.
// * void put(int key, int value) Update the value of the key if the key exists. 
//   Otherwise, add the key-value pair to the cache. 
//   If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.
//
// Runtime: 1053 ms, faster than 28.52% of Go online submissions for LRU Cache.
// Memory Usage: 73.6 MB, less than 75.90% of Go online submissions for LRU Cache.
//
type LRUCache struct {
    Capacity int
    Data map[int]int
    lru *LRUQueue
}

func Constructor(capacity int) LRUCache {
    return LRUCache{capacity, make(map[int]int), NewLRUQueue()}
}

func (this *LRUCache) Get(key int) int {
    if v, ok := this.Data[key]; ok {
        this.lru.Flush(key)
        return v
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int)  {
    this.Data[key] = value
    this.lru.Flush(key)

    if len(this.Data) > this.Capacity {
        if tail, ok := this.lru.Pop(); ok {
            delete(this.Data, tail)
        }
    }
}

type DoublyNode struct {
    Val int
    Next *DoublyNode
    Prev *DoublyNode
}

type LRUQueue struct {
    Head *DoublyNode
    Tail *DoublyNode
    keyMap map[int]*DoublyNode
}

func NewLRUQueue() *LRUQueue {
    return &LRUQueue{nil, nil, make(map[int]*DoublyNode)}
}

func (q *LRUQueue) Flush(val int) {
    if node, ok := q.keyMap[val]; !ok {
        node = &DoublyNode{val, nil, nil}
        node.Next = q.Head
        if q.Head != nil {
            q.Head.Prev = node
        }
        q.Head = node
        if q.Tail == nil {
            q.Tail = node
        }
        q.keyMap[val] = node
    } else {
        if node.Val == q.Head.Val {
            return
        }
        prev, next := node.Prev, node.Next
        prev.Next = next
        if next != nil {
            next.Prev = prev
        } else {
            q.Tail = prev
        }
        q.Head.Prev = node
        node.Next = q.Head
        q.Head = node
    }
}

func (q *LRUQueue) Pop() (val int, ok bool) {
    if q.Tail == nil {
        val, ok = -1, false
        return
    }

    val, ok = q.Tail.Val, true
    delete(q.keyMap, val)

    if q.Tail.Prev == nil {
        q.Head, q.Tail = nil, nil
        return
    }

    prev := q.Tail.Prev
    prev.Next = nil
    q.Tail = prev
    return
}
