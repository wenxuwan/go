package structTree

type treeNode struct {
	value       int
	left, right *treeNode
}

func Sort(values []int) {
	var root *treeNode
	for _, v := range values {
		root = build_tree(root, v)
	}
	//fmt.Printf("%v", cap(values2))
	MiddleTraverse(values[:0], root)
}

func build_tree(node *treeNode, value int) *treeNode {
	if node == nil {
		node = new(treeNode)
		node.value = value
		return node
	}
	if value < node.value {
		node.left = build_tree(node.left, value)
	} else {
		node.right = build_tree(node.right, value)
	}

	return node
}

func MiddleTraverse(values []int, node *treeNode) []int {
	if node != nil {
		values = MiddleTraverse(values, node.left)
		values = append(values, node.value)
		values = MiddleTraverse(values, node.right)
	}
	return values

}
