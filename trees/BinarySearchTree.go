package trees

type BinaryTreeNode struct {
	value       int
	left, right *BinaryTreeNode
}

func binary_tree_constructor(item int) BinaryTreeNode {
	return BinaryTreeNode{item, nil, nil}
}

func (n *BinaryTreeNode) addItem(item int) {
	node := &BinaryTreeNode{item, nil, nil}
  if n.value < item {
    if n.right == nil {
      n.right = node
    } else {
      n.right.addItem(item)
    }
  }else if n.value >= item {
    if n.left == nil {
      n.left = node
    } else {
      n.left.addItem(item)
    }
  }
}

// could implement find and delete
