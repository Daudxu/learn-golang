package main

import (
	"fmt"

	"github.com/nanmu42/etherscan-api"
)

func main() {
	// 创建连接指定网络的客户端
	client := etherscan.New(etherscan.Rinkby, "[W6127C4ADI6M4RBYV5NVXA4VJPTEJH5NYV]")
	
	// 或者，如果你要调用的是EtherScan家族的BscScan：
	//
	// client := etherscan.NewCustomized(etherscan.Customization{
	// Timeout:       15 * time.Second,
	// Key:           "You key here",
	// BaseURL:       "https://api.bscscan.com/api?",
	// Verbose:       false,
	// })	
	

	// 查询账户以太坊余额
	balance, err := client.AccountBalance("0x1537f0d523a264d3bBDf8d4A4e8778cd65b6D166")
	if err != nil {
		panic(err)
	}
	// 余额以 *big.Int 的类型呈现，单位为 wei
	fmt.Println(balance.Int())

	// // 查询token余额
	// tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")

	// // 查询出入指定地址的ERC20转账列表
	// transfers, err := client.ERC20Transfers("contractAddress", "address", startBlock, endBlock, page, offset)
}