package main

import "fmt"

func main() {}

const empty = '\000'

type (
	node struct {
		char byte
		next *node
	}

	trie struct {
		root *node
	}
)

func newNode() *node {
	return &node{char: empty}
}

func (t *trie) insert(word string) {
	current := t.root

	for _, c := range []byte(word) {

		if current.next == nil {
			current.next = newNode()
		}

		if current.next.char == empty {
			current.next.char = c
			return
		}

		if current.next.char != c {
			current.next = nil
			return
		}

		current = current.next
	}
}

func (t *trie) traverse() string {
	result := []byte{}

	current := t.root.next

	for current != nil {
		result = append(result, current.char)
		current = current.next
	}

	fmt.Printf("fot from traverse: %s\n", result)
	return string(result)
}

func newTrie(initial string) *trie {
	t := &trie{root: &node{}}
	t.insert(initial)
	return t
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Compare each character column by column
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			return strs[0][:i]
		}

		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
}

// longestCommonPrefix finds the longest common prefix among a slice of strings.
