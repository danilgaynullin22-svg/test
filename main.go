package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	randomTime := (rand.Intn(5)+1) * 100
	fmt.Println("Worker", i)

	time.Sleep( time.Duration(randomTime) * time.Millisecond)

	fmt.Println("Готово")
}
func main() {
	wg := &sync.WaitGroup{}
	for i := range 5{
		wg.Add(1)
		go worker(wg, i+1)
	}

	wg.Wait()
}