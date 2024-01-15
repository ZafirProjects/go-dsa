package trees

func dfs_bst_search(curr *BinaryTreeNode, needle int) bool {
  if curr == nil {
    return false
  }

  if curr.value == needle {
    return true
  }

  if curr.value < needle {
    return dfs_bst_search(curr.right, needle)
  }
  return dfs_bst_search(curr.left, needle)
}

func dfs_bst(head BinaryTreeNode, needle int) bool {
  return dfs_bst_search(&head, needle)
}
