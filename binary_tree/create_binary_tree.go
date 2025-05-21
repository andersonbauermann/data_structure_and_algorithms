package binary_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}
	return root
}

func Search(root *Node, value int) *Node {
	if root == nil || root.Value == value {
		return root
	}
	if value < root.Value {
		return Search(root.Left, value)
	}
	return Search(root.Right, value)
}

func PrintBinaryTreeSearch() {
	binaryTree := BinaryTree{}
	valuesToInsert := []int{10, 5, 15, 3, 7, 12, 18}

	for _, value := range valuesToInsert {
		binaryTree.Root = Insert(binaryTree.Root, value)
	}

	if Search(binaryTree.Root, 7) != nil {
		println("Valor encontrado na árvore binária.")
	} else {
		println("Valor não encontrado na árvore binária.")
	}
}
