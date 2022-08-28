package trie

type Node struct {
	Value    interface{}
	Children map[int32]*Node
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
	}
}
