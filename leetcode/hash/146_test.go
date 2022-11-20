package leetcode

import (
	"reflect"
	"testing"
)

// 146. LRU Cache
func TestLRUCache(t *testing.T) {
    var cache LRUCache
	t.Run("Test Constructor()", func(t *testing.T) {
        cache = Constructor(2)
        want := 2
        got := cache.Capacity
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Put()", func(t *testing.T) {
        cache.Put(1,1)
        cache.Put(2,2)
        want := map[int]int{
            1: 1,
            2: 2,
        }
        got := cache.Data
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Get()", func(t *testing.T) {
        want := 1
        got := cache.Get(1)
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Put()", func(t *testing.T) {
        cache.Put(3,3)
        want := map[int]int{
            1: 1,
            3: 3,
        }
        got := cache.Data
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Get()", func(t *testing.T) {
        want := -1
        got := cache.Get(2)
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Put()", func(t *testing.T) {
        cache.Put(4,4)
        want := map[int]int{
            3: 3,
            4: 4,
        }
        got := cache.Data
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Get()", func(t *testing.T) {
        inputs := []int{1,3,4}
        wants := []int{-1,3,4}
        for i, in := range inputs {
            got := cache.Get(in)
            want := wants[i]
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })
}
