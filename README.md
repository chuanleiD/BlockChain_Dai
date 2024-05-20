# README: 工作量证明的区块链原型系统

## 文件介绍

`project\BlockChain_Dai`：节点一，工程文件

`project\BlockChain_Dai2`：节点二，工程文件（拷贝后修改网络配置）

`project\BlockChain_Dai3`：节点三，工程文件（同二）

`pic\showdatabase.md`：实验运行中，三个节点的输出信息

<img src="pic\image-20240423182701783.png" alt="image-20240423182701783" style="zoom: 50%;" />

## 如何运行：

编译方法：（Windows系统下编译好的二进制文件在`bin`）

```
cd BlockChain_Dai
go build -o blockchain01.exe BlockChain_Dai

cd BlockChain_Dai2
go build -o blockchain02.exe BlockChain_Dai

cd BlockChain_Dai3
go build -o blockchain03.exe BlockChain_Dai
```

```go
// 项目结构
module BlockChain_Dai

go 1.21

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
```

如何运行：

1、分别运行`blockchain01.exe`、`blockchain02.exe`、`blockchain03.exe`

2、或尝试如下`bat`脚本

```bat
@echo off 
start cmd /k "blockchain01\blockchain01.exe"
start cmd /k "blockchain02\blockchain02.exe"
start cmd /k "blockchain03\blockchain03.exe"
```

3、或在GoLand等IDE下运行工程文件

`project\BlockChain_Dai`

`project\BlockChain_Dai2`

`project\BlockChain_Dai3`