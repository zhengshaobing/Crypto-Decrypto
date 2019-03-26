package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端 dial err：", err)
		return
	}
	//功能一：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入【终端】
	for {
		//从终端读取一行输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("终端读取失败，err ：", err)
			return
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出....")
			break
		}
		//再将读取的发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn Write err:", err)
		}
	}

}