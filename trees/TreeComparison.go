package trees

func compare_binary_tree(a, b *BinaryTreeNode) bool {
  if a == nil && b == nil {
    return true
  }

  if a == nil || b == nil {
    return false
  }

  if a.value != b.value {
    return false
  }

  return compare_binary_tree(a.left, b.left) && compare_binary_tree(a.right, b.right)
}
