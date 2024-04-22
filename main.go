package main

import (
	"fmt"
	"math/big"
)

// 全局变量，存储配置信息
var globalConfig Config
var TargetInt *big.Int

func init() {
	globalConfig.configGet("config.json")

	TargetInt = big.NewInt(1) //获取真实难度值的数值
	TargetInt.Lsh(TargetInt, uint(256-int32(globalConfig.Target)))

	globalConfig.show()
}

func main() {
	// 读取本地存储，恢复区块链
	Blockchain := new(BlockChain)
	Blockchain.ReadBlockchain()
	Blockchain.Show()

	err := Blockchain.SendBlockChain()
	if err != nil {
		fmt.Println("SendBlockChain error:", err)
		return
	}

	return
}

// main 函数
func main22() {
	// 读取本地存储，恢复区块链
	Blockchain := new(BlockChain)
	Blockchain.ReadBlockchain()
	Blockchain.Show()

	/*
		blockchainJSON := Blockchain.Serialize()
		Blockchain.Deserialize(blockchainJSON)
		fmt.Println("After Deserialize:")
		Blockchain.Show()
	*/

	// 基于当前区块链生成新区块
	Block01 := new(Block)
	Block01.init()
	Blockchain.AddPrevMessage(Block01) // 基于当前区块链补全信息
	Block01.show()

	// 挖矿
	result := false
	for result == false {
		result = Block01.RoundMine()
	}
	fmt.Println("finish mining")

	// 展现当前的区块链信息
	Blockchain.AddNewBlock(Block01)

	//持久化存储：
	Blockchain.SaveBlockchain()

	Blockchain.Show()

}
