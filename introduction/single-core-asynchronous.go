package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func rangeSumAsynchronous(start int, end int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sum := 0
	for num := start; num <= end; num++ {
		sum += num
	}
	return sum
}

func SingleCoreAsynchronous() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(3)

	start := time.Now()

	sum := 0
	go func() {
		sum += rangeSumAsynchronous(1, 10000, &wg)
	}()
	go func() {
		sum += rangeSumAsynchronous(10001, 20000, &wg)
	}()
	go func() {
		sum += rangeSumAsynchronous(20001, 40000, &wg)
	}()

	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Duration: %d microseconds\n", duration.Microseconds())
	fmt.Printf("Sum: %d\n", sum)
}
