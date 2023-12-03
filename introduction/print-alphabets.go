package main

import (
	"fmt"
	"runtime"
	"sync"
)

func lowerCase(goChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'a'; ch <= 'z'; ch++ {
		goChannel <- fmt.Sprintf("%c", ch)
	}
}

func upperCase(goChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'A'; ch <= 'Z'; ch++ {
		goChannel <- fmt.Sprintf("%c", ch)
	}
}

func PrintAlphabetsParallelism() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)
	goChannel := make(chan string)
	go lowerCase(goChannel, &wg)
	go upperCase(goChannel, &wg)

	go func() {
		for ch := range goChannel {
			fmt.Printf("%s", ch)
		}
	}()

	wg.Wait()
	fmt.Println()
}
