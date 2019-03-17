package indexes

import (
	"sort"
	"sync"
	"unicode"
)

type int64set map[int64]struct{}

// TrieNode defines the node fields for Trie
type TrieNode struct {
	key      rune
	values   int64set
	parent   *TrieNode
	children map[rune]*TrieNode
	mx       sync.RWMutex
}

// Trie data structure that stores
// keys of type string and values of type int64
type Trie struct {
	root *TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		key:      0,
		values:   int64set{},
		parent:   nil,
		children: map[rune]*TrieNode{},
		mx:       sync.RWMutex{},
	}
}

// NewTrie makes a reference to a new Trie
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// Add adds a given key/value pair to the trie
func (t *Trie) Add(key string, val int64) bool {
	curr := t.root
	curr.mx.Lock()
	defer curr.mx.Unlock()

	for _, runeVal := range key {
		// make a new node with the current runeValue if doens't exist
		if _, found := curr.children[runeVal]; !found {
			newNode := newTrieNode()
			newNode.parent = curr
			newNode.key = runeVal
			curr.children[runeVal] = newNode
		}
		curr = curr.children[runeVal]
	}

	// check if the value already exists in the trie, return false if yes
	if _, ok := curr.values[val]; ok {
		return false
	}
	// otherwise add the value as an entry to the targeet node
	curr.values[val] = struct{}{}
	return true
}

// ReturnPrefixMatches returns the first n values that match a given prefix string
func (t *Trie) ReturnPrefixMatches(n int, prefix string) []int64 {
	curr := t.root
	for _, runeVal := range prefix {
		curr = curr.children[runeVal]
	}
	var result []int64
	result = curr.searchTrie(n, result)
	return result
}

// changes set's keys to a slice
func keysToSlice(s int64set) []int64 {
	keys := make([]int64, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

func (curr *TrieNode) searchTrie(n int, result []int64) []int64 {
	if len(result) < n {
		valueSlice := keysToSlice(curr.values)
		for _, val := range valueSlice {
			result = append(result, val)
			if len(result) == n {
				break
			}
		}
		childrenKeys := make([]rune, 0, len(curr.children))
		for k := range curr.children {
			childrenKeys = append(childrenKeys, k)
		}
		sort.Slice(childrenKeys, func(i, j int) bool {
			return unicode.ToLower(childrenKeys[i]) < unicode.ToLower(childrenKeys[j])
		})
		for _, k := range childrenKeys {
			child := curr.children[k]
			return child.searchTrie(n, result)
		}
	}
	return result
}

// Delete deletes a given keyval pair from the trie
func (t *Trie) Delete(key string, val int64) {
	// t.root.mx.Lock()
	// defer t.root.mx.Unlock()
	runeSlice := []rune(key)
	length := len(runeSlice)
	t.root.deleteHelper(runeSlice, val, 0, length)
}

func (curr *TrieNode) deleteHelper(runes []rune, val int64, index int, originalLength int) {
	if originalLength == index { // find the original value to be deleted
		delete(curr.values, val)
		curr.deleteHelper(runes, val, -1, originalLength)
	} else if len(curr.children) == 0 && len(curr.values) == 0 { // if the current node is an empty leaf
		currentRune := runes[index]
		curr = curr.parent
		delete(curr.children, currentRune)
		curr.parent.deleteHelper(runes, val, -1, originalLength)
	} else if index == -1 { // if the target key/value has been deleted, and parent has content inside
		return
	} else if index < originalLength { // still traversing to the target key/value pair
		currentRune := runes[index]
		curr.children[currentRune].deleteHelper(runes, val, index+1, originalLength)
	}
}
