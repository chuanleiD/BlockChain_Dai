package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"
)

// BlockHead 区块头，用于哈希计算与验证
type BlockHead struct {
	Version    int32  // 版本号，4个字节，由配置文件指定
	PrevHash   string // 上一个区块的Hash值，哈希值是32个字节，encode为string时为64字节
	Time       int32  // 时间戳，4个字节，BlockHead.getTime()
	TargetBits int32  // 难度值，4个字节，表示左移的位数
	Nonce      int64  // 随机数，4个字节，BlockHead.getNonce()
	TxRoot     string // 交易根哈希，64个字节，BlockContain的Hash
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
}

// ------------------------------------------------------------------------
// 初始化区块，仍需补充PrevHash、TxRoot、Height信息
func (b *Block) init() {
	b.Head.Version = int32(globalConfig.Version) //int转int32
	b.Head.PrevHash = ""
	b.Head.getTime()
	b.Head.TargetBits = int32(globalConfig.Target) //int转int32
	b.Head.getNonce()
	b.Head.TxRoot = ""

	b.Contain.MinerId = globalConfig.MinerID //string
	b.Contain.Height = 0
	fmt.Println("[func (b *Block) init()]:")
	fmt.Println("Version:", b.Head.Version, ",Time:", b.Head.Time, ",TargetBits:", b.Head.TargetBits, ",Nonce:", b.Head.Nonce, ",MinerId:", b.Contain.MinerId)
}

// CreateGenesisBlock 创世块生成函数，注意：不同节点拥有同样的创世块
func CreateGenesisBlock(data string) *Block {
	fmt.Println("[func CreateGenesisBlock(data string) *Block]:")
	b := new(Block)
	b.Head.Version = int32(globalConfig.Version)                                         //int转int32
	b.Head.PrevHash = "0000000000000000000000000000000000000000000000000000000000000000" // 创世块的 PrevHash 是全零的哈希值
	b.Head.Time = 0
	b.Head.TargetBits = 29   //int转int32
	b.Head.Nonce = 347306002 //target为29时

	b.Contain.MinerId = data //string
	b.Contain.Height = 0     //创世块高度为0

	b.Head.TxRoot = b.getTxRoot() //TxRoot应该等于BlockContain的哈希值
	fmt.Println("CreateGenesisBlock success with data:", b.Contain.MinerId)
	return b
}

// 每次计算哈希值后，更新 Nonce
func (b *Block) update() {
	//b.Head.getTime() //每秒进行上百万次计算，时间戳不变
	b.Head.Nonce = b.Head.Nonce + 1
}

// 序列化区块头，用于哈希值计算
func (bh *BlockHead) serialize() []byte {
	bytes, err := json.Marshal(bh)
	if err != nil {
		log.Panic(err)
	}
	return bytes
}

// 获得区块头的哈希值，Byte格式，用于获取前一个块头的哈希值(作为key存储)
func (bh *BlockHead) getHashByte() []byte {
	hash := sha256.Sum256(bh.serialize())
	//存储时，将哈希值保存为[]byte
	hashByte := hash[:]
	return hashByte
}

// 获得区块头的哈希值，String格式，用于获取前一个块头的哈希值(块内存储)
func (bh *BlockHead) getHashString() string {
	hash := sha256.Sum256(bh.serialize())
	//存储时，将哈希值保存为字符串
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

// 获得区块头的哈希值，Big.Int格式，用于进行比较(工作量证明中进行比较)
func (bh *BlockHead) getHashInt() big.Int {
	hash := sha256.Sum256(bh.serialize()) // 获取当前块的哈希值
	hashBytes := hash[:]
	hashInt := big.Int{}
	hashInt.SetBytes(hashBytes)
	return hashInt
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
// 得到 BlockContain 后，计算 TxRoot 的哈希值，返回
func (b *Block) getTxRoot() string {
	// 将 BlockContain 结构体转换为 JSON 格式的字节数组
	data, err := json.Marshal(b.Contain)
	if err != nil {
		fmt.Println("Error marshaling BlockContain:", err)
		return ""
	}

	// 计算 SHA-256 哈希值
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])

	fmt.Println("TxRoot:", hashString)
	return hashString
}

// ------------------------------------------------------------------------
func (b *Block) show() {
	fmt.Println("[func (b *Block) show()]:")
	fmt.Println("Version:", b.Head.Version, ",PrevHash:", b.Head.PrevHash, ",Time:", b.Head.Time, ",TargetBits:", b.Head.TargetBits, ",Nonce:", b.Head.Nonce, ",TxRoot:", b.Head.TxRoot)
	fmt.Println("MinerId:", b.Contain.MinerId, ",Height:", b.Contain.Height)
}
