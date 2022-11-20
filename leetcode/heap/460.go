package leetcode

import (
        "container/heap"
        "fmt"
)

// 460. LFU Cache
//
// Design and implement a data structure for a Least Frequently Used (LFU) cache.
// Implement the LFUCache class:
// * LFUCache(int capacity) Initializes the object with the capacity of the data structure.
// * int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
// * void put(int key, int value) Update the value of the key if present, or inserts the key 
//   if not already present. When the cache reaches its capacity, it should invalidate and remove 
//   the least frequently used key before inserting a new item. For this problem, when there is a tie 
//   (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
// To determine the least frequently used key, a use counter is maintained for each key in the cache. 
// The key with the smallest use counter is the least frequently used key.
// When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). 
// The use counter for a key in the cache is incremented either a get or put operation is called on it.
// The functions get and put must each run in O(1) average time complexity.
//
// Runtime: 682 ms, faster than 90.00% of Go online submissions for LFU Cache.
// Memory Usage: 85.9 MB, less than 42.50% of Go online submissions for LFU Cache.
//
type LFUItem struct {
    d int
    c int
    k int
    v int
}

func (this LFUItem) String() string {
    return fmt.Sprintf("{d:%d,c:%d,k:%d,v:%d}", this.d, this.c, this.k, this.v)
}
func (this *LFUItem) Count() int { return this.c }
func (this *LFUItem) Flush() { this.c++ }
func (this *LFUItem) Data() int { return this.d }
func (this *LFUItem) SetData(data int) { this.d = data }
func (this *LFUItem) Key() int { return this.k }
func (this *LFUItem) Version() int { return this.v }
func (this *LFUItem) SetVersion(v int) { this.v = v }

type LFUData struct {
    data []*LFUItem
    indices map[int]int
    version int
}

func (this *LFUData) Index(key int) int { return this.indices[key] }
func (this *LFUData) SetIndex(key, idx int) { this.indices[key] = idx }

func (this LFUData) Len() int { return len(this.data) }
func (this LFUData) Less(i, j int) bool {
    if this.data[i].Count() != this.data[j].Count() {
        return this.data[i].Count() < this.data[j].Count()
    }
    return this.data[i].Version() < this.data[j].Version()
}
func (this LFUData) Swap(i, j int) {
    this.data[i], this.data[j] = this.data[j], this.data[i]
    this.SetIndex(this.data[i].Key(), i)
    this.SetIndex(this.data[j].Key(), j)
}

func (this *LFUData) Push(x interface{}) {
    v := x.(*LFUItem)
    this.data = append(this.data, v)
    this.indices[v.Key()] = len(this.data)-1
}

func (this *LFUData) Pop() interface{} {
    n := len(this.data)
    x := this.data[n-1]
    this.data = this.data[:n-1]
    delete(this.indices, x.Key())
    return x
}

func (this *LFUData) Add(key, val int) {
    //fmt.Printf("Add %d, version: %d, indices: %v\n", key, this.version, this.indices)
    this.version++
    item := &LFUItem{k: key, d: val, c: 1, v: this.version}
    heap.Push(this, item)
}

func (this *LFUData) Replace(key, val int) {
    //fmt.Printf("Replace %d, version: %d, indices: %v\n", key, this.version, this.indices)
    this.version++
    idx := this.indices[key]
    this.data[idx].SetData(val)
    this.data[idx].SetVersion(this.version)
    this.data[idx].Flush()
    heap.Fix(this, idx)
}

func (this *LFUData) Get(key int) int {
    idx, ok := this.indices[key]
    if !ok {
        return -1
    }
    //fmt.Printf("Search %d, version: %d, indices: %v\n", key, this.version, this.indices)
    this.version++
    v := this.data[idx].Data()
    this.data[idx].SetVersion(this.version)
    this.data[idx].Flush()
    heap.Fix(this, idx)
    return v
}

func (this *LFUData) Remove() {
    heap.Pop(this)
}

func (this *LFUData) HasKey(key int) bool {
    _, ok := this.indices[key]
    return ok
}

type LFUCache struct {
    data *LFUData
    c int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        data: &LFUData{data: []*LFUItem{}, indices: make(map[int]int)},
        c: capacity,
    }
}

func (this *LFUCache) Put(key int, value int) {
    if this.data.HasKey(key) {
        this.data.Replace(key, value)
        return
    }
    if this.data.Len() == this.c {
        this.data.Remove()
    }
    this.data.Add(key, value)
}

func (this *LFUCache) Get(key int) int {
    return this.data.Get(key)
}

