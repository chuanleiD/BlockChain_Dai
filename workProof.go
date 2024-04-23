package main

import (
	"fmt"
)

// Mine 检查区块是否合法，希望速度快，所以不加入fmt.Println()等调试信息
func (b *Block) Mine() (result bool) {
	i := 0
	result = false
	for i < 100000 {
		hashInt := b.Head.getHashInt()    // 获取区块头的哈希值
		if hashInt.Cmp(TargetInt) == -1 { //hashInt < target
			result = true
			break
		} else {
			b.update()
			i += 1
		}
	}
	return result
}

// RoundMine 检查区块是否合法，希望速度快，所以不加入fmt.Println()等调试信息
func (b *Block) RoundMine() (result bool) {
	//startTime := time.Now() //计时开始
	i := 0
	result = false
	for i < 10 {
		if b.Mine() == false { // 检查区块是否合法
			b.update() // 更新区块 Nonce+1
			i += 1
		} else {
			result = true
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Println("RoundMine success with Nonce:", b.Head.Nonce)
			hashInt := b.Head.getHashInt()
			fmt.Println("HashInt__:", hashInt.String())
			fmt.Println("targetInt:", TargetInt.String())
			fmt.Println("---------------------------------------------------------------------------------------")
			break
		}
	}
	return result
}
