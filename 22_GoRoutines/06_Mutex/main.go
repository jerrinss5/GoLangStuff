package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// getting the lock for this block of code
		mutex.Lock()

		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter: ", counter)

		mutex.Unlock()
		// releasing the lock for other threads
	}
	wg.Done()
}
