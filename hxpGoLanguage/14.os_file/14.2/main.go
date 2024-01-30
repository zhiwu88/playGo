package main

import (
	"fmt"
	"os"
)

// 测试文件的打开和关闭
func main() {

	// 打开文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	// 打印file，可以看出file是个指针
	fmt.Printf("file=%v", file)

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
