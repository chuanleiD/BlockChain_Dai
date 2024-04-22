package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// SendBlockChain 参考POW共识协议内容，节点间传输整个区块信息
func (bc *BlockChain) SendBlockChain() error {
	// 序列化区块链
	blockchainJSON := bc.Serialize()

	// 连接到程序2的地址和端口
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting:", err)
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Error Close:", err)
		}
	}(conn)

	// 发送序列化后的 BlockChain 字节流
	_, err = conn.Write(blockchainJSON)
	if err != nil {
		log.Fatal("Error sending:", err)
		return err
	}

	log.Println("BlockChain sent successfully")
	return nil
}

// Listen 监听本地端口，接收其他节点发送的区块链
func Listen() {
	fmt.Println("Listening at localhost:8080...")
	// 监听本地端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			log.Fatal("Error Listen Close:", err)
		}
	}(ln)

	for {
		// 等待连接
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error accepting:", err)
		}

		// 开启一个协程来处理连接
		go handleConnection(conn)
	}
}

// handleConnection 处理接收到的连接
func handleConnection(conn net.Conn) {
	fmt.Println("Handling new connection...")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Error handleConnection Close:", err)
		}
	}(conn)

	// 创建一个字节切片，用于存储接收到的数据
	var receivedData []byte

	// 读取接收到的数据
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading:", err)
			break
		}

		// 将读取到的数据追加到 receivedData 中
		receivedData = append(receivedData, buffer[:n]...)

		// 如果读取到的数据量小于缓冲区大小，说明数据已经接收完毕
		if n < len(buffer) {
			break
		}
	}

	// 反序列化 BlockChain
	receivedBlockchain := new(BlockChain)
	err := json.Unmarshal(receivedData, receivedBlockchain)
	if err != nil {
		log.Println("Error deserializing:", err)
		return
	}

	fmt.Println("Received BlockChain:")
	receivedBlockchain.Show()
	blockBuffer = append(blockBuffer, receivedBlockchain)
}
