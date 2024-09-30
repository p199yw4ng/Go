package main

import (
	"fmt"
)

func main() {
	var n, T int
	fmt.Scan(&n, &T)
	task := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(task[i])
	}
	fmt.Println(T)
}
