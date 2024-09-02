package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("Enter first tree array: ")
	var arr1, arr2 [][]int
	json.Unmarshal([]byte(readStringInput()), &arr1)

	fmt.Print("Enter second tree array: ")
	json.Unmarshal([]byte(readStringInput()), &arr2)

	finalState := minimumDiameterAfterMerge(arr1, arr2)

	fmt.Print("Result: ", finalState, "\n")
}

func readStringInput() (val string) {
	_, err := fmt.Scanf("%s", &val)

	if err != nil {
		panic(err)
	}

	return
}

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	diameter1 := findDiameter(edges1)
	diameter2 := findDiameter(edges2)

	answer := getRadius(diameter1) + 1 + getRadius(diameter2)
	answer = max(answer, diameter1)
	answer = max(answer, diameter2)

	return answer
}

func findDiameter(edges [][]int) int {
	n := len(edges) + 1
	conn := connections(n, edges)
	farthest, _ := bfs(n, conn, 0)
	_, diameter := bfs(n, conn, farthest)
	return diameter
}

func getRadius(diameter int) int {
	return diameter - diameter/2
}

func connections(n int, edges [][]int) [][]int {
	conn := make([][]int, n)
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		conn[node1] = append(conn[node1], node2)
		conn[node2] = append(conn[node2], node1)
	}

	return conn
}

func bfs(n int, conn [][]int, startNode int) (farthestNode int, maxHeight int) {
	height := 0
	item := []int{startNode, height}
	queue := [][]int{item}
	visited := make([]bool, n)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		node, height := item[0], item[1]

		if visited[node] {
			continue
		}
		visited[node] = true

		if height > maxHeight {
			maxHeight = height
			farthestNode = node
		}

		for _, neighbor := range conn[node] {
			if !visited[neighbor] {
				queue = append(queue, []int{neighbor, height + 1})
			}
		}
	}

	return farthestNode, maxHeight
}
