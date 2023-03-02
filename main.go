package main

import (
	"fmt"
)

func main() {
	n, prereq := 1, [][]int{}
	order := findOrder(n, prereq)
	if len(order) == 0 {
		fmt.Println("No valid course order found.")
	} else {
		fmt.Println("Course order:", order)
	}
}

func findOrder(n int, prereq [][]int) []int {
	if n == 0 {
		return []int{}
	}
	var conn [][]int = make([][]int, n)
	for _, p := range prereq {
		conn[p[0]] = append(conn[p[0]], p[1])
	}
	var visited []int = make([]int, n)
	var order []int
	for i := 0; i < n; i++ {
		if visited[i] != 0 {
			continue
		}
		if !dfs(conn, i, visited, &order) {
			return []int{}
		}
	}
	return order
}

func dfs(conn [][]int, curr int, visited []int, order *[]int) bool {
	if visited[curr] == -1 {
		return false
	}
	visited[curr] = -1
	for _, c := range conn[curr] {
		if visited[c] == 1 {
			continue
		}
		if visited[c] == 2 {
			return false
		}
		if !dfs(conn, c, visited, order) {
			return false
		}
	}
	visited[curr] = 1
	*order = append(*order, curr)
	return true
}
