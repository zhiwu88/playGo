package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:05:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

/* 测试方法，开2个终端，执行任意以下命令测试。结果显示两个终端都能输出，因为使用了go关键字，已支持多并发。
nc localhost 8000
curl --http0.9 localhost:8000
go run netcat1/main.go
*/
