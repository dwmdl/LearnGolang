package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	sumCh := make(chan int, 3)
	numGoroutines := 3
	chunkSize := len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		go sumChunk(sumCh, arr[start:end])
	}

	totalSum := 0
	for range numGoroutines {
		totalSum += <-sumCh
	}

	fmt.Println(totalSum)
}

func sumChunk(ch chan int, arr []int) {
	var res int
	for _, num := range arr {
		res += num
	}

	ch <- res
}
