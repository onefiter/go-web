//go:build v1

package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	// 使用WaitGroup等待 10 个 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()
			// 对变量 count 执行 10 次加 1
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	// 等待 10 个 goroutine完成
	wg.Wait()
	fmt.Println(count)
}
