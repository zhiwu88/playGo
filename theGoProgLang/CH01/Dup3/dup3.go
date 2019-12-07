//go run dup3.go 1.txt
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//map 存储一个键/值对集合
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//ReadFile函数返回一个可以转化为字符串的字节slice
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//fmt.Println(data)
		//示例：fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))  //  ["foo" "bar" "baz"]
		//解析：string(data)就是byte切片转为string，使用换行符\n作为分割符，换行的ASCII码是10。
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	//对map的迭代每次能迭代出两个结果，map里一个元素对应的键和值。顺序是随机的。
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
