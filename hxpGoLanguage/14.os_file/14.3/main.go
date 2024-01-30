package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件的内容并显示在终端（带缓冲区的方式），使用 os.Open 、file.Close 、bufio.NewReader() 、reader.ReadString
func main() {

	// 打开文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	// 当函数退出时，及时关闭file文件描述符，避免内存溢出
	defer file.Close()

	// 创建一个 *Reader ，默认缓冲区是4096字节（defaultBufSize = 4096）
	// 带缓冲区的好处就是并不是一次性把整个文件读取到内存，而是读取一部分处理一部分，可以处理较大的文件。
	reader := bufio.NewReader(file)
	// 循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 每次循环读取到一个换行符就结束
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		// 测试输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
