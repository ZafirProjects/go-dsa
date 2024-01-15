package trees

func dfs_search(curr *BinaryTreeNode, needle int) bool {
  if curr == nil {
    return false
  }

  if curr.value == needle {
    return true
  }

  if curr.value < needle {
    return dfs_search(curr.right, needle)
  }

  return dfs_search(curr.left, needle)
}

func dfs(head BinaryTreeNode, needle int) bool {
  return dfs_search(&head, needle)
}
