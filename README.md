# README: 工作量证明的区块链原型系统

[TOC]

## 一、项目需求：

利用(Python\Go\Rust)等语言实现一个**工作量证明的区块链原型系统**，验证工作量证明区块链的生成过程。相对于真实的比特币系统可以做如下简化：

- 可以通过配置IP地址指定连接的节点，无需动态加入。
- 无需模拟用户交易，共识节点模拟生成固定大小的块即可（区块头大小固定）。
- 出块速度无需自动调整，通过配置文件设置难度即可。

## 二、项目原理：

[工作量证明的区块链原型系统——项目原理](项目原理.md)

[工作量证明的区块链原型系统——关键细节介绍](关键细节介绍.md)

## 三、文件介绍

`BlockChain_Dai`：单节点工程文件。

`blockchain.db` ：当前本地的节点状态。

`config.json`：配置文件：

```json
{
  "Version":1,
  "Target":29,
  "MinerId":"Miner01", // ID
  "RoundTime":1000000 
}
```

项目结构：

```
// 项目结构
module BlockChain_Dai

go 1.21

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
```

<img src="pic\image-20240423182701783.png" alt="image-20240423182701783" style="zoom: 50%;" /> 

## 四、如何运行：

编译方法：

```
cd BlockChain_Dai
go build -o blockchain01.exe BlockChain_Dai
```

单区块链节点的文件结构：

<img src="D:\0go_work\Go_WorkSpace\BlockChain_Dai\pic\PixPin_2024-08-16_18-10-57.png" alt="PixPin_2024-08-16_18-10-57" style="zoom: 67%;" /> 

如何运行：

创建编译后的三份副本，详见：`\bin`：

1、分别运行`blockchain01.exe`、`blockchain02.exe`、`blockchain03.exe`

2、或在bin文件夹下尝试如下`bat`脚本

```bat
@echo off 
start cmd /k "blockchain01\blockchain01.exe"
start cmd /k "blockchain02\blockchain02.exe"
start cmd /k "blockchain03\blockchain03.exe"
```
