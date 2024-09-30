package main

import (
	"fmt"
)

func main() {

	var n, m int

	n = 4
	m = 13

	reward := []int{0, 1111, 525, 1031, 55, 0, 0, 722, 0, 430, 1221, 29, 711}
	var dp = make([][][]int, n+1)

	// 每层的状态3^n, 初始3^0 = 1
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, pow(3, n))
		for j := 0; j < pow(3, n); j++ {
			dp[i][j] = make([]int, m+1)
		}

	}
	dp[0][0][m] = 0

	for i := 1; i <= n; i++ {

		var a1, a2, a3 int
		var b1, b2, b3 int
		fmt.Scan(&a1, &a2, &a3)
		fmt.Scan(&b1, &b2, &b3)
		var a = []int{a1, a2, a3}
		var b = []int{b1, b2, b3}

		for j := 0; j < pow(3, i); j++ {
			//复制上一层状态
			copy(dp[i][j], dp[i-1][j/3])

			dp[i][j][m] += a[j%3]

			dp[i][j][b[j%3]-1]++
		}

	}

	// fmt.Println(dp)
	var max = 0
	for i := 0; i < pow(3, n); i++ {
		for j := 0; j < m; j++ {
			if dp[n][i][j] >= 3 {
				dp[n][i][m] += reward[j]
			}
		}
		if dp[n][i][m] > max {
			max = dp[n][i][m]
		}

	}
	fmt.Println(max)

}

func pow(n, m int) int {
	sum := 1
	for i := 0; i < m; i++ {
		sum *= n
	}
	return sum
}

// 3 5 3
// 3 2 4
// 2 3 7
// 2 11 5
// 4 0 6
// 10 2 13
// 10 5 196
// 1 12 8
