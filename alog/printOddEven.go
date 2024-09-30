package alog

import (
	"fmt"
	"sync"
)

// 打印奇数和偶数
func printOdds(wg *sync.WaitGroup, oddsChan, evensChan chan bool, num int) {
	defer wg.Done()
	for i := 1; i <= num; i += 2 {
		<-evensChan
		fmt.Println(i)
		if i < num {
			oddsChan <- true
		}
	}
}
func printEvens(wg *sync.WaitGroup, oddsChan, evensChan chan bool, num int) {
	defer wg.Done()
	for i := 0; i <= num; i += 2 {
		<-oddsChan
		fmt.Println(i)
		if i < num {
			evensChan <- true
		}
	}
}

func PrintOddAndEven(num int) {
	var wg sync.WaitGroup
	odds, evens := make(chan bool), make(chan bool)
	wg.Add(2)
	go printOdds(&wg, odds, evens, num)
	go printEvens(&wg, odds, evens, num)
	odds <- true
	wg.Wait()
}
