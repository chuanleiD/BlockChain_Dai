package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "Mr.Dai"

func (b *Block) Serialize() []byte {
	byteBlock, err := json.Marshal(b)
	if err != nil {
		log.Panic(err)
	}
	return byteBlock
}

func (b *Block) Deserialize(byteBlock []byte) {
	err := json.Unmarshal(byteBlock, b)
	if err != nil {
		fmt.Printf("json.Marshal,err:%s", err)
	}
}

// ReadBlockchain 读取本地数据库中的区块链到内存中
func (bc *BlockChain) ReadBlockchain() {
	fmt.Println("[func (bc *Blockchain) ReadBlockchain()]:")
	// 打开一个 BoltDB 文件 "blockchain.db"，如果不存在则创建一个
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db) // 在函数退出时关闭数据库

	// 读写的形式打开数据库中的区块链，blocksBucket = "blocks"
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		// 如果数据库中不存在区块链就创建一个
		if bucket == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")

			// 创建创世块，并写入内存
			genesis := CreateGenesisBlock(genesisCoinbaseData) // 写入创世块
			bc.AddNewBlock(genesis)                            // 将创世块添加到区块链中

			b, err := tx.CreateBucket([]byte(blocksBucket)) // 创建一个新的 bucket
			if err != nil {
				log.Panic(err)
			}
			// 存储创世块，key = genesis.Head.Time，value = genesis.Serialize()
			err = b.Put(Int32ToBytes(genesis.Head.Time), genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
		} else { // 如果数据库中已经存在区块链，就读取出来
			cursor := bucket.Cursor()
			for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
				fmt.Printf("Key: %d, Value: %s\n", BytesToInt32(key), value)
				block := new(Block)
				block.Deserialize(value)
				bc.AddNewBlock(block)
			}

		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return
}

// SaveBlockchain 将内存中的区块链保存到本地数据库中
func (bc *BlockChain) SaveBlockchain() {
	fmt.Println("[func (bc *Blockchain) SaveBlockchain()]:")
	// 打开一个 BoltDB 文件 "blockchain.db"，如果不存在则创建一个
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db) // 在函数退出时关闭数据库

	// 读写的形式打开数据库中的区块链，blocksBucket = "blocks"
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket)) // 获取指定的桶

		if bucket != nil { // 如果桶存在
			// 删除桶
			err := tx.DeleteBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
		}

		// 重新创建一个同名的空桶
		_, err = tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}
		bucket = tx.Bucket([]byte(blocksBucket))

		// 将区块链中的所有区块存入数据库
		for _, block := range bc.Blocks {
			err = bucket.Put(Int32ToBytes(block.Head.Time), block.Serialize())
			fmt.Println("Save block:", block.Head.Time, block.Serialize())
			if err != nil {
				log.Panic(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return
}

// ------------------------------------------------------------------------
// 为持久化存储，定义两个类型转换函数

// Int32ToBytes 将int32转换为[]byte
func Int32ToBytes(num int32) []byte {
	timeBytes, err := json.Marshal(num)
	if err != nil {
		log.Panic(err)
	}
	return timeBytes
}

// BytesToInt32 将[]byte转换为int32
func BytesToInt32(data []byte) int32 {
	var num int32
	err := json.Unmarshal(data, &num)
	if err != nil {
		log.Panic(err)
	}
	return num
}

// ------------------------------------------------------------------------
