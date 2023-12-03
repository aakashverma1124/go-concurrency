package main

import "fmt"

func main() {

	fmt.Println("Single Core Synchronously")
	SingleCoreSynchronous()
	fmt.Println("--------------------------------------")
	fmt.Println("Single Core Asynchronously")
	SingleCoreAsynchronous()
	fmt.Println("--------------------------------------")
	fmt.Println("Multi Core Synchronously")
	MultiCoreSynchronous()
	fmt.Println("--------------------------------------")
	fmt.Println("Multi Core Asynchronously")
	MultiCoreAsynchronous()
	fmt.Println("--------------------------------------")
	fmt.Println("Parallelism")
	ParallelismDriverFunction()
	fmt.Println("--------------------------------------")
}
