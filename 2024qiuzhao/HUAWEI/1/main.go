package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	//最大深度

	//读取边

	//
	neiberMap := make(map[int][]int, 0)

	for i := 0; i < n; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		//添加边
		neiberMap[u] = append(neiberMap[u], v)
		neiberMap[v] = append(neiberMap[v], u)
	}

	//是否访问过
	visit := make([]bool, n)
	// fmt.Println(neiberMap)
	fmt.Println(dfs(neiberMap, m, 0, k, visit) - 1)

	//求取边数<=k的个数

}

// dfs 查询
func dfs(neiberMap map[int][]int, user, curDepth, maxDepth int, visit []bool) int {
	if visit[user] || curDepth > maxDepth {
		return 0
	}
	visit[user] = true
	// 进来就有+1了
	count := 1

	for _, neiber := range neiberMap[user] {
		count += dfs(neiberMap, neiber, curDepth+1, maxDepth, visit)
	}
	return count

}
