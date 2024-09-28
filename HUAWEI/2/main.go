package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	//读取边
	martrix := make([][]bool, n)

	for i := 0; i < n; i++ {
		martrix[i] = make([]bool, n)
	}
	//读取不能放置的位置，1代表有障碍
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		martrix[x][y] = true
	}

	fmt.Println(search(martrix, n))
	//求取边数<=k的个数

}

// 搜索算法
func search(martrix [][]bool, n int) int {
	var ans int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if full(martrix, i, j, n) {
				martrix[i][j] = true
				martrix[i+1][j] = true
				martrix[i][j+1] = true
				martrix[i+1][j+1] = true
				ans++
			}
		}
	}

	return ans
}

// 能否填充
func full(martrix [][]bool, i, j, n int) bool {
	//边界无法填充

	if i+1 >= n || j+1 >= n {
		return false
	}
	for k := 0; k < 2; k++ {
		for l := 0; l < 2; l++ {
			if martrix[i+k][j+l] {
				return false
			}
		}
	}

	return true
}
