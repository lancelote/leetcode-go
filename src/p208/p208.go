package p208

type Node struct {
	children map[rune]*Node
	end      bool
}

func newNode() *Node {
	return &Node{children: make(map[rune]*Node), end: false}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: newNode()}
}

func (this *Trie) Insert(word string) {
	node := this.root

	for _, x := range word {
		nextNode, ok := node.children[x]
		if !ok {
			nextNode = newNode()
			node.children[x] = nextNode
		}
		node = nextNode
	}

	node.end = true
}

func (this *Trie) Search(word string) bool {
	node := this.root

	for _, x := range word {
		nextNode, ok := node.children[x]
		if !ok {
			return false
		}

		node = nextNode
	}

	return node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	for _, x := range prefix {
		nextNode, ok := node.children[x]
		if !ok {
			return false
		}

		node = nextNode
	}

	return true
}
