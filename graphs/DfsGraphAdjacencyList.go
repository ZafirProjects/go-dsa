package graphs

type GraphEdge struct {
	to, weight int
}

type WeightedAdjacencyList [][]GraphEdge

func walk_dfs_graph(graph WeightedAdjacencyList, curr, needle int, seen []bool) bool {

	if seen[curr] {
		return false
	}

	seen[curr] = true

	// pre
	path = append(path, curr)
	if curr == needle {
		return true
	}

	// recurse
	list := graph[curr]
	for _, v := range list {
		if walk_dfs_graph(graph, v.to, needle, seen) {
			return true
		}
	}

	// post
	path = path[:len(path)-1]

	return false
}

var path []int

func dfs_graph_adjacency_list(graph WeightedAdjacencyList, start, needle int) []int {
	var seen []bool
	for i := range seen {
		seen[i] = false
	}

	walk_dfs_graph(graph, start, needle, seen)
	return path
}
