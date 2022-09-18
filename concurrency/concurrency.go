package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sleepConcurrentTask()
		fmt.Printf("done %v time\n", i)
	}
	wg.Wait()
}

func sleepConcurrentTask() {
	fmt.Println("going to sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("waking up")
	wg.Done()
}
