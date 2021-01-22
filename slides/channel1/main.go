package main

import (
	"fmt"
	"strconv"
	"time"
)

// START OMIT
func worker(workQueue chan int, workerNumber int) {
	for work, isOpen := <-workQueue; isOpen; work, isOpen = <-workQueue {
		fmt.Println(strconv.Itoa(workerNumber) + " got work: " + strconv.Itoa(work))
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Channel closed")
}
func sender(workQueue chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("send work " + strconv.Itoa(i))
		workQueue <- i
	}
	fmt.Println("close channel")
	close(workQueue)
}
func main() {
	workQueue := make(chan int, 10)
	go sender(workQueue)
	go worker(workQueue, 1)
	go worker(workQueue, 2)
	worker(workQueue, 3)
}

// END OMIT
