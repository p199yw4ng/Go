package main

import (
	"fmt"
)

var xorToRoot []int
var xorCount map[int]int
var graph [][][2]int

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	// 初始化表，
	graph = make([][][2]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([][2]int, 0)
	}

	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		graph[u] = append(graph[u], [2]int{v, w})
		graph[v] = append(graph[v], [2]int{u, w})
	}

	//初始化抑或
	xorToRoot = make([]int, n+1)
	xorCount = make(map[int]int)

	dfs(1, -1, 0)

	//查询
	for i := 0; i < q; i++ {
		var u, k int
		fmt.Scan(&u, &k)

		target := xorToRoot[u] ^ k
		fmt.Println(xorCount[target])

	}

}

func dfs(node, parent int, xor int) {
	xorToRoot[node] = xor
	xorCount[xor]++

	for _, edge := range graph[node] {
		nb := int(edge[0])
		weight := edge[1]

		if nb != parent {
			dfs(nb, node, xor^weight)
		}
	}
}
