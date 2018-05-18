package trie

// Node 节点
type Node struct {
	char  rune
	child map[rune]*Node
	Data  interface{}
	final bool
	Depth int
}

// Trie 树状结构
type Trie struct {
	root *Node
}

func newNode() *Node {
	return &Node{
		child: make(map[rune]*Node),
	}
}

// NewTrie init a new trie
func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

// Add 向树中添加数据
func (t *Trie) Add(text string, data interface{}) {
	node := t.root
	for _, v := range text {
		ret, ok := node.child[v]
		if !ok {
			ret = newNode()
			ret.char = v
			ret.Depth = node.Depth + 1
			node.child[v] = ret
		}
		node = ret
	}
	node.final = true
	node.Data = data
}

func (t *Trie) findNode(text string) (result *Node) {
	node := t.root

	for _, v := range text {
		ret, ok := node.child[v]
		if !ok {
			return
		}
		node = ret
	}
	result = node
	return
}

func (t *Trie) colloction(n *Node) (nodes []*Node) {
	if n == nil {
		return
	}

	if n.final && len(n.child) == 0 {
		nodes = append(nodes, n)
		return
	}

	queue := make([]*Node, 0, 10)
	queue = append(queue, n)

	for i := 0; i < len(queue); i++ {
		if queue[i].final {
			nodes = append(nodes, queue[i])
			if len(queue[i].child) == 0 {
				continue
			}
		}

		for _, v := range queue[i].child {
			queue = append(queue, v)
		}
	}
	return
}

// PrefixSearch 前缀搜索
func (t *Trie) PrefixSearch(text string) (nodes []*Node) {
	resNode := t.findNode(text)
	if resNode == nil {
		return
	}
	nodes = t.colloction(resNode)
	return
}

// BadWord 敏感词替换
func (t *Trie) BadWord(text string, replace ...string) (result string, find bool) {
	node := t.root

	var char = []rune(text)
	var left []rune
	var start int

	replaceStr := "*"
	if len(replace) != 0 {
		replaceStr = replace[0]
	}

	for index, v := range char {
		ret, ok := node.child[v]
		if !ok {
			node = t.root
			ret, ok := node.child[v]
			if ok {
				left = append(left, char[start:index]...)
				start = index
				node = ret

				if index == len(char)-1 {
					left = append(left, char[start:index+1]...)
				}
				continue
			}

			left = append(left, char[start:index+1]...)
			start = index + 1

			continue
		}

		node = ret

		if ret.final {
			rs := replaceStr
			for i := 0; i < len(char[start:index]); i++ {
				rs += replaceStr
			}
			left = append(left, []rune(rs)...)
			start = index + 1
			find = true
		}

		if index == len(char)-1 {
			left = append(left, char[start:index+1]...)
		}
	}

	result = string(left)
	return
}
