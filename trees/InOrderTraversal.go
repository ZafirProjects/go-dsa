package trees

var inpath = []int{}

func inwalk(curr *BinaryTreeNode) []int {
	if curr == nil {
		return inpath
	}
	// pre
	// recurse
	inwalk(curr.left)
  inpath = append(inpath, curr.value)
	inwalk(curr.right)
	// post
	return inpath
}

func in_order_search(head *BinaryTreeNode) []int {
	return inwalk(head)
}
