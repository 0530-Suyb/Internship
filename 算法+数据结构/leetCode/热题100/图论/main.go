package main

import (
	"fmt"
	"math"
)

func dijkstra(edges [][]int, n int, k int) []int {
	m := make([][][]int, n)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1:])
		m[edge[1]] = append(m[edge[1]], []int{edge[0], edge[2]})
	}

	visited := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[k] = 0

	for {
		minN := -1
		for i := 0; i < n; i++ {
			if !visited[i] && (minN == -1 || dist[i] < dist[minN]) {
				minN = i
			}
		}
		if minN == -1 || dist[minN] == math.MaxInt {
			break
		}
		visited[minN] = true
		for _, edge := range m[minN] {
			dist[edge[0]] = min(dist[edge[0]], dist[minN]+edge[1])
		}
	}

	return dist
}

func main() {
	fmt.Println(dijkstra([][]int{
		{0, 1, 1},
		{0, 2, 3},
		{0, 3, 6},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}, 4, 0))
}
