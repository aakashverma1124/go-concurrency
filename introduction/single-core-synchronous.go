package main

import (
	"fmt"
	"runtime"
	"time"
)

func rangeSumSynchronous(start int, end int) int {
	sum := 0
	for num := start; num <= end; num++ {
		sum += num
	}
	return sum
}

func SingleCoreSynchronous() {
	runtime.GOMAXPROCS(1)
	start := time.Now()
	sum := 0
	sum += rangeSumSynchronous(1, 10000)
	sum += rangeSumSynchronous(10001, 20000)
	sum += rangeSumSynchronous(20001, 40000)
	duration := time.Since(start)
	fmt.Printf("Duration: %d microseconds\n", duration.Microseconds())
	fmt.Printf("Sum: %d\n", sum)
}
