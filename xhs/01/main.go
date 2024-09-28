package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var zan = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&zan[i])
	}
	var ans int

	ans = count(zan, k)
	fmt.Println(ans)
}

func count(zan []int, k int) int {
	zanmap := make(map[int]int)

	ans := 0

	for _, v := range zan {
		tar := v ^ k
		if _, ok := zanmap[tar]; ok {
			ans += zanmap[tar]
		}
		zanmap[v]++
	}
	return ans
}
