package wkzip

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readFile() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"} // 替换为你要读取的txt文件列表
	var wg sync.WaitGroup
	wg.Add(len(files))
	lines := make(chan string) // 用于接收读取的行数
	for _, file := range files {
		go func(f string) {
			defer wg.Done()
			file, err := os.Open(f) // 打开文件
			if err != nil {
				fmt.Printf("打开文件 %s 时发生错误: %v\n", f, err)
				return
			}
			defer file.Close()
			scanner := bufio.NewScanner(file) // 创建Scanner对象
			for scanner.Scan() {              // 逐行扫描文件内容
				lines <- scanner.Text() // 将读取的行数发送到结果通道
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("读取文件 %s 时发生错误: %v\n", f, err)
			}
		}(file)
	}
	go func() {
		wg.Wait()    // 等待所有协程完成
		close(lines) // 关闭结果通道
	}()
	// 收集并打印所有
	for line := range lines {
		fmt.Println("读取的行数:", line)
	}
}
