type TreeNode struct {
	key int
	val int
	l   *TreeNode
	r   *TreeNode
}

type TreeMap struct {
	root *TreeNode
}

func NewTreeMap() *TreeMap {
	return &TreeMap{}
}

func insertNode(node *TreeNode, key, val int) *TreeNode {
	if node == nil {
		return &TreeNode{key: key, val: val}
	}
	if key < node.key {
		node.l = insertNode(node.l, key, val)
	} else if key > node.key {
		node.r = insertNode(node.r, key, val)
	} else {
		node.val = val
	}
	return node
}

func (tm *TreeMap) Insert(key, val int) {
	tm.root = insertNode(tm.root, key, val)
}

func (tm *TreeMap) Get(key int) int {
	curr := tm.root
	for curr != nil {
		if key < curr.key {
			curr = curr.l
		} else if key > curr.key {
			curr = curr.r
		} else {
			return curr.val
		}
	}
	return -1
}

func (tm *TreeMap) GetMin() int {
	if tm.root == nil {
		return -1
	}
	curr := tm.root
	for curr.l != nil {
		curr = curr.l
	}
	return curr.val
}

func (tm *TreeMap) GetMax() int {
	if tm.root == nil {
		return -1
	}
	curr := tm.root
	for curr.r != nil {
		curr = curr.r
	}
	return curr.val
}

func minNode(node *TreeNode) *TreeNode {
	for node.l != nil {
		node = node.l
	}
	return node
}

func removeNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.l = removeNode(node.l, key)
	} else if key > node.key {
		node.r = removeNode(node.r, key)
	} else {
		if node.l == nil {
			return node.r
		} else if node.r == nil {
			return node.l
		}
		successor := minNode(node.r)
		node.key = successor.key
		node.val = successor.val
		node.r = removeNode(node.r, successor.key)
	}
	return node
}

func (tm *TreeMap) Remove(key int) {
	tm.root = removeNode(tm.root, key)
}

func inorder(node *TreeNode, keys *[]int) {
	if node == nil {
		return
	}
	inorder(node.l, keys)
	*keys = append(*keys, node.key)
	inorder(node.r, keys)
}

func (tm *TreeMap) GetInorderKeys() []int {
	keys := []int{}
	inorder(tm.root, &keys)
	return keys
}