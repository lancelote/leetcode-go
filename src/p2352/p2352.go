package p2352

type TrieNode struct {
	count    int
	children map[int]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[int]*TrieNode),
	}
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) insert(s []int) {
	node := t.root
	for _, x := range s {
		if _, ok := node.children[x]; !ok {
			node.children[x] = newTrieNode()
		}
		node = node.children[x]
	}
	node.count++
}

func (t *Trie) search(s []int) int {
	node := t.root
	for _, x := range s {
		if _, ok := node.children[x]; !ok {
			return 0
		}
		node = node.children[x]
	}
	return node.count
}

func equalPairs(grid [][]int) int {
	count := 0
	trie := newTrie()
	for _, row := range grid {
		trie.insert(row)
	}

	for c := 0; c < len(grid[0]); c++ {
		col := make([]int, 0, len(grid))
		for r := 0; r < len(grid); r++ {
			col = append(col, grid[r][c])
		}
		count += trie.search(col)
	}

	return count
}
