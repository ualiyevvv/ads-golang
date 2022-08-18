package adsgraphs

var visited []int

func OneGraph(slice []int, startPoint int, count int) int {
	visited = append(visited, startPoint)
	for i, n := range slice {
		if n != 0 && !isVisitedGraph(i) {
			count = OneGraph(GlobalArr[i], i, count)
		}
	}
	count++
	return count
}

func isVisitedGraph(n int) bool {
	for _, r := range visited {
		if r == n {
			return true
		}
	}
	return false
}
