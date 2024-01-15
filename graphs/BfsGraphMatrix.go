package graphs

type WeightedAdjacencyMatrix [][]int

func bfs_graph_matrix(graph WeightedAdjacencyMatrix, start, needle int) []int {
	seen := make([]bool, len(graph))
	for i := range seen {
		seen[i] = false
	}
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	q := []int{start}

	for len(q) != 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := range adjs {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
	}

	curr := needle
	var out []int

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) != 0 {
		return append([]int{start}, reverse(out)...)
	}

	return []int{}
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
