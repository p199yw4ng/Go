package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var volt = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&volt[i])
	}
	sort.Ints(volt)
	var sum, ans int
	for _, v := range volt {

		if v+sum <= m {
			ans++
			sum += v
		}
	}
	fmt.Println(ans)
}
