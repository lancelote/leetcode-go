package p1268

type Node struct {
	children []*Node
	end      bool
}

func newNode() *Node {
	return &Node{children: make([]*Node, 26)}
}

type Trie struct {
	root *Node
}

func newTrie() Trie {
	return Trie{root: newNode()}
}

func (t *Trie) insert(word string) {
	node := t.root

	for _, x := range word {
		nextNode := node.children[x-'a']
		if nextNode == nil {
			nextNode = newNode()
			node.children[x-'a'] = nextNode
		}
		node = nextNode
	}

	node.end = true
}

func (t *Trie) threeByPrefix(prefix string) []string {
	r := []string{}

	stack := []rune{}
	node := t.root
	for _, x := range prefix {
		nextNode := node.children[x-'a']
		if nextNode == nil {
			return r
		}
		stack = append(stack, x)
		node = nextNode
	}

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if len(r) == 3 {
			return
		}

		if n == nil {
			return
		}

		if n.end {
			r = append(r, string(stack))
		}

		for i, nextNode := range n.children {
			stack = append(stack, rune('a'+i))
			dfs(nextNode)
			stack = stack[0 : len(stack)-1]
		}
	}

	dfs(node)
	return r
}

func suggestedProducts(products []string, searchWord string) [][]string {
	r := [][]string{}
	t := newTrie()
	for _, p := range products {
		t.insert(p)
	}

	for i := 0; i < len(searchWord); i++ {
		r = append(r, t.threeByPrefix(searchWord[0:i+1]))
	}

	return r
}
