package adsgraphs

import "fmt"

func isValidGraph(arr [][]int) bool {
	degreeOfVertexes := 0
	countOfEdge := 0
	visited := []int{}
	for i := 0; i < len(arr); i++ {
		visited = append(visited, i)
		for j := 0; j < len(arr); j++ {
			if !isVisited(visited, j) && arr[i][j] != 0 {
				countOfEdge++
			}
			if arr[i][j] == 1 {
				degreeOfVertexes++
			} else if arr[i][j] == 2 {
				degreeOfVertexes += 2
			}
		}
	}

	fmt.Println(arr)

	countOfEdge *= 2
	return countOfEdge == degreeOfVertexes
}

func isVisited(visited []int, item int) bool {
	for _, r := range visited {
		if r == item {
			return true
		}
	}

	return false
}
