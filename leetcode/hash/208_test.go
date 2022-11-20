package leetcode

import (
	"testing"
)

// 208. Implement Trie (Prefix Tree)
func TestTrie(t *testing.T) {
    var trie Trie
    /*
	t.Run("Test all()", func(t *testing.T) {
        trie = Constructor3()
        inputs := []string{"p", "pr", "pre", "pref"}
        wants := [][]bool{
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
        }
        for i, in := range inputs {
            want := wants[i]
            trie.Insert(in)
            got := []bool{trie.Search(in), trie.StartsWith(in)}
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })
    */


	t.Run("Test insert all()", func(t *testing.T) {
        trie = Constructor3()
        inputs := []string{"app", "apple", "beer", "add", "jam", "rental"}
        /*
        wants := [][]bool{
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
            []bool{true, true},
        }
        */
        trie.StartsWith("a")
        for _, in := range inputs {
            trie.Insert(in)
            /*
            want := wants[i]
            got := []bool{trie.Search(in), trie.StartsWith(in)}
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
            */
        }
    })
	t.Run("Test search all()", func(t *testing.T) {
        inputs := []string{"apps","app","ad","applepie","rest","jan","rent","beer","jam"}
        wants := []bool{false,true,false,false,false,false,false,true,true}
        for i, in := range inputs {
            want := wants[i]
            got := trie.Search(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

	t.Run("Test Constructor()", func(t *testing.T) {
        trie = Constructor3()
        want := false
        got := false
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Insert()", func(t *testing.T) {
        trie.Insert("ab")
        want := false
        got := trie.Search("abc")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Search()", func(t *testing.T) {
        want := true
        got := trie.Search("ab")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test StartsWith()", func(t *testing.T) {
        want := false
        got := trie.StartsWith("abc")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test StartsWith()", func(t *testing.T) {
        want := true
        got := trie.StartsWith("ab")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Insert()", func(t *testing.T) {
        want := true
        trie.Insert("ab")
        trie.Insert("abc")
        got := trie.Search("abc")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
        got = trie.StartsWith("abc")
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
}
