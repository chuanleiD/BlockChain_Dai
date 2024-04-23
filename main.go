package main

import (
	"fmt"
	"math/big"
)

// 全局变量，存储配置信息
var globalConfig Config
var TargetInt *big.Int
var blockBuffer []*BlockChain

func init() {
	globalConfig.configGet("config.json")

	TargetInt = big.NewInt(1) //获取真实难度值的数值
	TargetInt.Lsh(TargetInt, uint(256-int32(globalConfig.Target)))

	globalConfig.show()
}

// main 函数
func main() {
	go Listen() // 监听网络

	// 读取本地存储，恢复区块链
	BlockchainSave := new(BlockChain)
	BlockchainSave.ReadBlockchain()
	BlockchainSave.Show()

	// 基于当前区块链生成新区块
	Block01 := BlockchainSave.NewBlock()

	// 挖矿
	for {
		result := Block01.RoundMine()
		if result == true { // 若本矿工发现了新区块
			// 展现当前的区块链信息
			BlockchainSave.AddNewBlock(Block01)
			//持久化存储：
			BlockchainSave.SaveBlockchain()
			fmt.Println("Block01.RoundMine() == true")
			Block01 = BlockchainSave.NewBlock() // 基于当前区块链生成新区块
			// 发送给其他矿工

			err := BlockchainSave.SendBlockChain("localhost:12002")
			if err != nil {
				fmt.Println("SendBlockChain error:", err)
				continue
			}
			err2 := BlockchainSave.SendBlockChain("localhost:12003")
			if err2 != nil {
				fmt.Println("SendBlockChain error:", err2)
				continue
			}
		} else {
			if len(blockBuffer) > 0 { // 若其他矿工发下了新区块
				for i := 0; i < len(blockBuffer); i++ {
					Blockchain := blockBuffer[i] // 若其他矿工的新区块链要更长
					if BlockchainSave.Blocks[len(BlockchainSave.Blocks)-1].Contain.Height < Blockchain.Blocks[len(Blockchain.Blocks)-1].Contain.Height {
						checkBufferBlock := Blockchain.Validate()
						if checkBufferBlock == true {
							fmt.Println("checkBufferBlock == true")
							BlockchainSave = Blockchain
							BlockchainSave.SaveBlockchain()
							BlockchainSave.Show()
						}
					}
				}
				blockBuffer = nil                   // 清空缓存
				Block01 = BlockchainSave.NewBlock() // 基于当前区块链生成新区块
			}
		}

	}

}

func main2() {
	// 读取本地存储，恢复区块链
	BlockchainSave := new(BlockChain)
	BlockchainSave.ReadBlockchain()
	BlockchainSave.Show()
}
