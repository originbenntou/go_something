package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 10)
	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(printv int) {
			log.Println(printv)
			time.Sleep(1000 * time.Millisecond)
			<- ch
			wg.Done()
		}(i)
	}
	wg.Wait()
}
