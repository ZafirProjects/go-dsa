package graphs

import (
	"math"
)

func dijksta_list(start, end int, graph WeightedAdjacencyList) []int {
	var seen []bool
	for i := range seen {
		seen[i] = false
	}

	var prev []int
	for i := range prev {
		prev[i] = -1
	}

	var dists []int
	for i := range dists {
		dists[i] = int(math.Inf(1))
	}

	dists[start] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	var out []int
	curr := end

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, start)
	return reverse(out)
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, v := range seen {
		if !v && dists[i] < int(math.Inf(1)) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := int(math.Inf(1))

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}

	return idx
}
