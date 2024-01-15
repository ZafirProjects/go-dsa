package trees

var postpath = []int{}

func postwalk(curr *BinaryTreeNode) []int {
	if curr == nil {
		return postpath
	}
	// pre
	// recurse
	postwalk(curr.left)
	postwalk(curr.right)
	// post
  postpath = append(postpath, curr.value)
	return postpath
}

func post_order_search(head *BinaryTreeNode) []int {
	return postwalk(head)
}
