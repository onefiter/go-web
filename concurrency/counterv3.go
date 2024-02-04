//go:build v3

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	wg.Add(10)

	// 启动10个goroutine

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加 10 万次
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count)

}
