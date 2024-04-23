package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
// 已知区块的结构
type Block struct {
	Head    BlockHead
	Contain BlockContain
	Parent  *Block
	Child   *[]Block
}
*/

// Blockchain 参考POW共识协议内容，节点间传输整个区块信息
type BlockChain struct {
	Blocks []*Block
}

// AddPrevMessage 为新区块添加前一个区块的信息
func (bc *BlockChain) AddPrevMessage(newBlock *Block) {
	fmt.Println("[func (bc *BlockChain) AddPrevMessage()]:")
	prevBlock := bc.Blocks[len(bc.Blocks)-1] // 获取上一个区块

	newBlock.Contain.Height = prevBlock.Contain.Height + 1 // 计算新区块的高度
	fmt.Println("newBlock.Contain.Height:", newBlock.Contain.Height)

	newBlock.Head.TxRoot = newBlock.getTxRoot() // 获取上一个区块的哈希值
	fmt.Println("newBlock.Head.TxRoot:", newBlock.Head.TxRoot)

	newBlock.Head.PrevHash = prevBlock.Head.getHashString() // 获取上一个区块的哈希值
	fmt.Println("newBlock.Head.PrevHash:", newBlock.Head.PrevHash)
}

// AddNewBlock 添加新区块
func (bc *BlockChain) AddNewBlock(newBlock *Block) {
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Show 显示区块链信息
func (bc *BlockChain) Show() {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("【Show the entire blockchain】:")
	for _, b := range bc.Blocks {
		b.show()
	}
	fmt.Println("---------------------------------------------------------------------------------------")
}

//---------------------------------------------------------------------
// 区块链整体的序列化与反序列化

// Serialize 使用方法：blockchainJSON := Blockchain.Serialize()
func (bc *BlockChain) Serialize() []byte {
	// 将 BlockChain 结构体序列化为 JSON 格式
	fmt.Println("[func (bc *BlockChain) Serialize()]:")
	blockchainJSON, err := json.Marshal(bc)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("blockchainJSON:", string(blockchainJSON))
	return blockchainJSON
}

// Deserialize 使用方法：Blockchain2 := new(BlockChain), Blockchain2.Deserialize(blockchainJSON)
func (bc *BlockChain) Deserialize(blockchainJSON []byte) {
	fmt.Println("[func (bc *BlockChain) Deserialize(blockchainJSON []byte)]:")
	err := json.Unmarshal(blockchainJSON, bc)
	if err != nil {
		fmt.Println("Failed to deserialize BlockChain:", err)
		return
	}
	fmt.Println("Deserialize success")
}

func (bc *BlockChain) NewBlock() (b *Block) {
	fmt.Println("[func (bc *BlockChain) NewBlock()]:")
	Block01 := new(Block)
	Block01.init()
	bc.AddPrevMessage(Block01) // 基于当前区块链补全信息
	Block01.show()
	return Block01
}

//---------------------------------------------------------------------
//

func (bc *BlockChain) Validate() bool {
	fmt.Println("[func (bc *BlockChain) Validate()]:")
	for i := len(bc.Blocks) - 1; i > 0; i-- {
		if bc.Blocks[i].Head.PrevHash != bc.Blocks[i-1].Head.getHashString() { // 比较当前区块的 PrevHash 与上一个区块的哈希值
			fmt.Println("BlockChain PrevHash is invalid!")
			return false
		}
		if bc.Blocks[i].Head.TxRoot != bc.Blocks[i].getTxRoot() { // 比较当前区块的 TxRoot 与区块内交易的哈希值
			fmt.Println("BlockChain TxRoot is invalid!")
			return false
		}
		hashInt := bc.Blocks[i].Head.getHashInt()
		if hashInt.Cmp(TargetInt) > -1 { // 比较当前区块的哈希值与目标值
			fmt.Println("BlockChain Target is invalid!")
			return false
		}
	}
	fmt.Println("BlockChain is valid!")
	return true
}
