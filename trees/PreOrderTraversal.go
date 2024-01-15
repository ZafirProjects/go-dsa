package trees

var prepath = []int{}

func prewalk(curr *BinaryTreeNode) []int {
	if curr == nil {
		return prepath
	}
	// pre
	prepath = append(prepath, curr.value)
	// recurse
	prewalk(curr.left)
	prewalk(curr.right)
	// post
	return prepath
}

func pre_order_search(head *BinaryTreeNode) []int {
	return prewalk(head)
}
