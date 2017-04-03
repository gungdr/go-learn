package dijsktra

import "fmt"

const MAX_INT int = 9999
const V int = 9

func minDistance(dist [V]int, sptSet [V]bool) int {
	min := MAX_INT
	var min_index int
	for v := 0; v < V; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			min_index = v
		}
	}
	return min_index
}

func printSolution(dist [V]int) {
	for i := 0; i < V; i++ {
		fmt.Println(i, dist[i])
	}
}
func dijkstra(graph [][]int, src int) {

	var dist [V]int
	var sptSet [V]bool

	//set all to infinite, and sptSet false
	for i := 0; i < V; i++ {
		dist[i] = MAX_INT
		sptSet[i] = false
	}

	dist[src] = 0

	for c := 0; c < V-1; c++ {

		min_dist := minDistance(dist, sptSet)
		sptSet[min_dist] = true
		for v := 0; v < V; v++ {
			if !sptSet[v] && graph[min_dist][v] > 0 &&
				dist[min_dist] != MAX_INT &&
				dist[min_dist]+graph[min_dist][v] < dist[v] {

				dist[v] = dist[min_dist] + graph[min_dist][v]
			}
		}

	}

	printSolution(dist)

}
func Generate() {
	graph := [][]int{{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0}}
	dijkstra(graph, 0)

}
