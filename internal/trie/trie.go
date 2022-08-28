package trie

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

func (t *Trie) Insert(key string, value interface{}) {
	cur := t.Root

	for _, ch := range key {
		node, ok := cur.Children[ch]
		if !ok {
			node = NewNode()
			cur.Children[ch] = node
		}

		cur = cur.Children[ch]
	}

	cur.Value = value
}

// Search returns the maximum match
func (t *Trie) Search(text string) (interface{}, bool) {
	cur := t.Root

	for _, ch := range text {
		node, ok := cur.Children[ch]
		if !ok {
			break
		}

		cur = node

		if len(cur.Children) == 0 {
			break
		}
	}

	if cur.Value != nil {
		return cur.Value, true
	}

	return nil, false
}

func (t *Trie) MaxDepth() int {
	return maxDepth(t.Root) - 1
}

func maxDepth(node *Node) int {
	max := 0

	for _, child := range node.Children {
		if v := maxDepth(child); v > max {
			max = v
		}
	}

	return max + 1
}
