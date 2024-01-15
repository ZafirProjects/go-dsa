package trees

func breadth_first_search(head BinaryTreeNode, needle int) bool {
  var q = []BinaryTreeNode{head}
  for len(q) != 0 {
    curr := q[0]
    q = q[1:]

    if curr.value == needle {
      return true
    }

    if curr.left != nil {
      q = append(q, *curr.left)
    }

    if curr.right != nil {
      q = append(q, *curr.right)
    }

  }
  return false
}
