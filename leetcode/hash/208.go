package leetcode

// 208. Implement Trie (Prefix Tree)
//
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store 
// and retrieve keys in a dataset of strings. There are various applications of this data structure, 
// such as autocomplete and spellchecker.
// Implement the Trie class:
// * Trie() Initializes the trie object.
// * void insert(String word) Inserts the string word into the trie.
// * boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), 
//   and false otherwise.
// * boolean startsWith(String prefix) Returns true if there is a previously inserted string word 
//   that has the prefix prefix, and false otherwise.
//
// Runtime: 57 ms, faster than 94.05% of Go online submissions for Implement Trie (Prefix Tree).
// Memory Usage: 18.6 MB, less than 27.50% of Go online submissions for Implement Trie (Prefix Tree).
//
type TrieNode struct {
    links []*TrieNode
    isEnd bool
}

type Trie struct {
    root *TrieNode
}

func NewTrieNode() *TrieNode {
    return &TrieNode{make([]*TrieNode, 26), false}
}

func (this *TrieNode) IsEnd() bool {
    return this.isEnd
}

func (this *TrieNode) SetEnd() {
    this.isEnd = true
}

func (this *TrieNode) ContainsKey(ch byte) bool {
    return this.links[ch-'a'] != nil
}

func (this *TrieNode) Get(ch byte) *TrieNode {
    return this.links[ch-'a']
}

func (this *TrieNode) Put(ch byte, node *TrieNode) {
    this.links[ch-'a'] = node
}

func Constructor3() Trie {
    return Trie{NewTrieNode()}
}

func (this *Trie) Insert(word string) {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if !node.ContainsKey(ch) {
            node.Put(ch, NewTrieNode())
        }
        node = node.Get(ch)
    }
    node.SetEnd()
}

func(this *Trie)searchPrefix(word string) *TrieNode {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if node.ContainsKey(ch) {
            node = node.Get(ch)
        } else {
            return nil
        }
    }
    return node
}

func (this *Trie) Search(word string) bool {
    node := this.searchPrefix(word)
    return node != nil && node.IsEnd()
}

func (this *Trie) StartsWith(prefix string) bool {
    node := this.searchPrefix(prefix)
    return node != nil
}

/*
type Trie struct {
    term string
    isEnd bool
    children map[byte]*Trie
}

func Constructor3() Trie {
    return Trie{"", false, make(map[byte]*Trie, 26)}
}

func (this *Trie) Print(indent int) {
    indentStr := ""
    for i := 0; i < indent; i++ {
        indentStr += " "
    }
    fmt.Printf("%s%v\n", indentStr, *this)
    for _, child := range this.children {
        child.Print(indent+4)
    }
}

func (this *Trie) match(word string) int {
    c := 0
    for c < len(this.term) && c < len(word) && this.term[c] == word[c] {
        c++
    }
    return c
}

func (this *Trie) Insert(word string) {
    pnode := this
    pos := 0
    for pos < len(word) {
        matched := pnode.match(word[pos:])
        if matched != len(pnode.term) {
            branch := Constructor3()
            branch.term = pnode.term[matched:]
            branch.isEnd = pnode.isEnd
            branch.children = pnode.children
            pnode.children = make(map[byte]*Trie, 26)
            pnode.children[branch.term[0]] = &branch
            pnode.term = pnode.term[:matched]
            pnode.isEnd = false
        }
        pos += matched
        if pos == len(word) {
            pnode.isEnd = true
            return
        }
        if child, ok := this.children[word[pos]]; ok {
            pnode = child
        } else {
            break
        }
    }
    node := Constructor3()
    node.term = word[pos:]
    node.isEnd = true
    pnode.children[word[pos]] = &node
}

func (this *Trie) Search(word string) bool {
    pnode, pos := this, 0
    for pos < len(word) {
        matched := pnode.match(word[pos:])
        if matched < len(pnode.term) {
            return false
        }
        pos += matched
        if pos == len(word) {
            if pnode.isEnd && matched == len(pnode.term) {
                return true
            }
            return false
        }
        if child, ok := pnode.children[word[pos]]; ok {
            pnode = child
        } else {
            return false
        }
    }
    return false
}

func (this *Trie) StartsWith(prefix string) bool {
    fmt.Printf("Prefix %s\n", prefix)
    fmt.Println("============== Trie ================")
    this.Print(0)
    pos, pnode := 0, this
    for pos < len(prefix) {
        matched := pnode.match(prefix[pos:])
        pos += matched
        if pos == len(prefix) {
            return true
        }
        if matched < len(pnode.term) {
            return false
        }
        if child, ok := pnode.children[prefix[pos]]; ok {
            pnode = child
        } else {
            return false
        }
    }
    return false
}
*/
