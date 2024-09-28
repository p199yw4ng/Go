package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {

		var n int
		fmt.Scan(&n)

		// 创建长度n的数组用来情况
		// idx 为 歌曲位置，值为0,1,-1
		var scores = make([]int, n)

		for j := 0; j < n; j++ {
			var x, y int
			fmt.Scan(&x, &y)
			scores[y-1] = x

		}

		fmt.Println(rerank(scores))

	}

}

func rerank(scores []int) string {
	ans := 0

	flag := false
	for _, v := range scores {
		ans += v
		// 排名先出现为-1 直接返回
		if ans >= 0 {
			if ans >= 1 {
				flag = true
			}
			continue
		} else {
			if !flag {
				return "NO"
			}

		}

	}

	// 逆序 若先出现 1 也直接返回
	flag = false
	ans = 0
	for i, _ := range scores {
		v := scores[len(scores)-i-1]
		ans += v
		// 排名第最后若为 1 直接返回
		if ans <= 0 {
			if ans <= -1 {
				flag = true
			}
			continue
		} else {
			if !flag {
				return "NO"
			}

		}

	}

	return "YES"

}
