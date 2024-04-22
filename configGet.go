package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config 配置文件
type Config struct {
	Version   int    `json:"Version"`
	Target    int    `json:"Target"`
	MinerID   string `json:"MinerId"`
	RoundTime int    `json:"RoundTime"`
}

func (c *Config) configGet(jsonUrl string) bool {
	// 读取 JSON 文件
	data, err := os.ReadFile(jsonUrl)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	// 解析 JSON 数据
	err = json.Unmarshal(data, c)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return false
	}

	return true
}

func (c *Config) show() {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("[func (c *Config) show()]:")
	fmt.Println("Version:", c.Version, ",Target:", c.Target, ",MinerID:", c.MinerID, ",RoundTime:", c.RoundTime)
	fmt.Println("TargetInt:", TargetInt)
	fmt.Println("---------------------------------------------------------------------------------------")
}

/*
// 配置文件读取方式
config01 := new(Config)
config01.configGet("config.json")

// 输出解析后的内容及其类型
fmt.Printf("Version: %d, Type: %T\n", config01.Version, config01.Version)
fmt.Printf("Target: %d, Type: %T\n", config01.Target, config01.Target)
fmt.Printf("Miner ID: %s, Type: %T\n", config01.MinerID, config01.MinerID)
fmt.Printf("Round Time: %d, Type: %T\n", config01.RoundTime, config01.RoundTime)
*/
