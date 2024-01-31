package wkzip

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Stat(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一个map来存储结果
	wordMap := make([]map[string]string, 0)
	// 创建一个scanner来读取文件内容
	scanner := bufio.NewScanner(file)
	titleNames := []string{}

	var lineNum int = 0
	for scanner.Scan() {

		words := strings.Fields(scanner.Text())
		if lineNum == 0 { // 首行为title
			titleNames = append(titleNames, words...)
			lineNum++

		} else {
			temp := make(map[string]string)
			for i := 0; i < len(words); i++ {

				temp[titleNames[i]] = strings.TrimSpace(words[i])

			}
			wordMap = append(wordMap, temp)
			lineNum++

		}

		// fmt.Printf("%T\n", words)
		// fmt.Printf("%d\n", len(words))
		// fmt.Println(words)
		// for _, word := range words {
		// 	fmt.Println(words)
		// 	// 去除单词两边的空格并转换为小写，然后增加计数器
		// 	word = strings.ToLower(strings.TrimSpace(word))
		// 	wordCount[word]++
		// }
	}
	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	// // 打印结果
	// for word, count := range wordCount {
	// 	fmt.Printf("%s: %d\n", word, count)
	// }

	fmt.Println(wordMap)
}
