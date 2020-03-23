package main

import (
	"fmt"
	"github.com/google/gops/agent"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	var a, b, c int
	wg := new(sync.WaitGroup)
	for {
		wg.Add(1)
		go func() {
			fmt.Println("博多")
			a++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("TEST", runtime.NumGoroutine())

	fmt.Println(a,b,c)
	time.Sleep(time.Hour)
}
