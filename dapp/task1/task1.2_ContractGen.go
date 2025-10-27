package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"homework/dapp/task1/myCounter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。

	具体任务

编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/
func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/hKtjTkxwP7TTEEJjwt8CV")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xaF112f3E479c02fcB6ea0E771725d0a0F8c95225")

	//加载合约
	myCounter, err := myCounter.NewMyCounter(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("b439f9e75ff2813c6bb0f56235edbd3eebe06c1b6f64c8f441a955dd1e0084c7")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := myCounter.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx.Hash() ", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	value, err := myCounter.GetCount(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("myCounter value: ", value)

}
