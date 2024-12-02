package main

import (
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
	time.Sleep(1 * time.Second)
	fmt.Println("bye")
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(wg)
	}
	wg.Wait()
	fmt.Println("end")
}
