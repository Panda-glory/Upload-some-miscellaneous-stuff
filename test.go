package main

import (
	"fmt"
	"net"
)

func QueryAvailablePort() uint16 {
	// 定义查询的端口范围
	startPort := 1024
	endPort := 49151

	// 逐个检查端口是否可用
	for port := startPort; port <= endPort; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			defer func(listener net.Listener) {
				err := listener.Close()
				if err != nil {
					return
				}
			}(listener)
			return uint16(port)
		}
	}

	// 如果没有可用端口，则返回0
	return 0
}

func main() {
	// 调用接口查询可用端口并打印结果
	port := QueryAvailablePort()
	if port != 0 {
		fmt.Printf("可用端口: %d\n", port)
	} else {
		fmt.Println("未找到可用端口")
	}
}
