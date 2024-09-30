package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x, y, z, k int
		fmt.Scan(&x, &y, &z, &k)
		fmt.Println(maxCount(x, y, z, k))
	}

}

func maxCount(x, y, z, k int) int {

	maxNum := 0
	//遍历 不同长宽高的桌子
	for a := 1; a*a <= k; a++ {
		if k%a == 0 {
			k1 := k / a
			for b := 1; b*b <= k1; b++ {
				if k1%b == 0 {
					c := k1 / b
					if b <= c {
						maxNum = max(maxNum, calMax(x, y, z, a, b, c))
					}
				}
			}
		}
	}
	return maxNum
}
func calMax(x, y, z, a, b, c int) int {

	way := 0

	// 不同方向上的摆放
	if a <= x && b <= y && c <= z {
		way = max(way, (x-a+1)*(y-b+1)*(z-c+1))
	}

	if a <= x && b <= z && c <= y {
		way = max(way, (x-a+1)*(y-c+1)*(z-b+1))
	}

	if a <= y && b <= x && c <= z {
		way = max(way, (x-b+1)*(y-a+1)*(z-c+1))
	}
	if a <= y && b <= z && c <= x {
		way = max(way, (x-c+1)*(y-a+1)*(z-b+1))
	}
	if a <= z && b <= x && c <= y {
		way = max(way, (z-a+1)*(x-b+1)*(y-c+1))
	}
	if a <= z && b <= y && c <= x {
		way = max(way, (z-a+1)*(y-b+1)*(x-c+1))
	}
	return way
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
