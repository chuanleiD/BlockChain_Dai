package main

import (
	"math/rand"
	"time"
)

// BlockHead 区块头，用于哈希计算与验证
type BlockHead struct {
	Version    int32  // 版本号，4个字节，由配置文件指定
	PrevHash   string // 上一个区块的Hash值，32个字节
	Time       int32  // 时间戳，4个字节，BlockHead.getTime()
	TargetBits uint32 // 难度值，4个字节，表示左移的位数
	Nonce      int64  // 随机数，4个字节，BlockHead.getNonce()
	TxRoot     string // 交易根哈希，32个字节，BlockContain的Hash
}

// BlockContain 区块内容，用于存储具体交易信息
type BlockContain struct {
	MinerId string // 矿工 ID，由配置文件指定
	Height  uint   // 区块高度，由上一个区块的BlockContain.Height+1得到
}

// Block 区块
type Block struct {
	Head    BlockHead
	Contain BlockContain
	Parent  *Block
	Child   *[]Block
}

// ------------------------------------------------------------------------

func (b *Block) init() {

	config01 := new(Config)
	config01.configGet("config.json")

	b.Head.Version = 1
	b.Head.PrevHash = ""
	b.Head.getTime()
	b.Head.TargetBits = 31
	b.Head.getNonce()
	b.Head.TxRoot = ""
	b.Contain.MinerId = "Miner"
	b.Contain.Height = 0
	b.Parent = nil
	b.Child = nil
}

// ------------------------------------------------------------------------

// 获得32位时间戳，每次区块链更新之后，更新一次就可以
func (bh *BlockHead) getTime() {
	currentTime := time.Now()
	bh.Time = int32(currentTime.Unix())
	return
}

// 获得32位随机数，每次区块链更新之后，更新一次就可以，随后每次挖矿时nonce = (nonce + 1) % math.MaxInt64
func (bh *BlockHead) getNonce() {
	bh.Nonce = rand.Int63()
	return
}

// ------------------------------------------------------------------------
