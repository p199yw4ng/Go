package main

import (
	"fmt"
)

func maxPathCost(matrix [][]int, n int) int {
	// 创建一个三维切片用于动态规划，dp[i][j][k] 表示第 i 步，第一个玩家选择第 j 行，第二个玩家选择第 k 行的最大路径成本
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	// 初始化第一步
	dp[1][0][0] = matrix[0][0]
	dp[1][0][1] = matrix[1][0]
	dp[1][1][0] = matrix[1][0]
	dp[1][1][1] = matrix[0][0]

	for i := 2; i <= n; i++ {
		// 第一个玩家选择第一行
		dp[i][0][0] = max(dp[i-1][1][0], dp[i-1][1][1]) + matrix[0][i-1]
		// 第一个玩家选择第二行
		dp[i][1][1] = max(dp[i-1][0][0], dp[i-1][0][1]) + matrix[1][i-1]
		// 第二个玩家选择第一行
		dp[i][0][1] = min(dp[i-1][1][0], dp[i-1][1][1]) + matrix[0][i-1]
		// 第二个玩家选择第二行
		dp[i][1][0] = min(dp[i-1][0][0], dp[i-1][0][1]) + matrix[1][i-1]
	}
	fmt.Println(dp)
	return max(dp[n][0][0], dp[n][1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//获取N
	var N int
	fmt.Scan(&N)

	matrix := make([][]int, 2)

	//获取两行元素
	var x int
	var row []int
	for i := 0; i < N; i++ {

		fmt.Scan(&x)
		row = append(row, x)
	}
	matrix[0] = row
	row = []int{}
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		row = append(row, x)
	}
	matrix[1] = row

	// fmt.Println(matrix)
	result := maxPathCost(matrix, N)
	fmt.Println("两个玩家都采取最优策略下的路径成本为：", result)
}
