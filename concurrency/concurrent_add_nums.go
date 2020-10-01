package main

import (
	"sync"
)

func add(start int, end int) int{
	sum := 0
	for i := start; i<=end; i++{
		sum += i
	}
	return sum
}
var wg sync.WaitGroup
func AddConcurrent(start int, end int, skip int) int {
	bufferedChan := make(chan int, 10000)
	i:= start

	for i<=end{
		foundSkip := skip
		if i + skip > end{
			foundSkip = end - i
		}
		wg.Add(1)
		go func(k int, l int) {
			bufferedChan <- add(k, k+l)
			wg.Done()
		}(i,foundSkip)
		i += foundSkip + 1
	}

	wg.Wait()
	close(bufferedChan)

	totalSum := 0
	for semi := range bufferedChan{
		//semi := <- bufferedChan
		totalSum += semi
	}

	return totalSum
}