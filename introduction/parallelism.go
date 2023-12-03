package main

import (
	"fmt"
	"runtime"
	"sync"
)

func valueSet1(goChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		goChannel <- i
	}
}

func valueSet2(goChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 5; i < 10; i++ {
		goChannel <- i
	}
}

func valueSet3(goChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i < 15; i++ {
		goChannel <- i
	}
}

func ParallelismDriverFunction() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(3)
	goChannel := make(chan int)

	go valueSet1(goChannel, &wg)
	go valueSet3(goChannel, &wg)
	go valueSet2(goChannel, &wg)

	for i := 0; i < 15; i++ {
		fmt.Print(<-goChannel, " ")
	}

	fmt.Println()

	wg.Wait()

}
