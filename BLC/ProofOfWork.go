package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 256位Hash里面前面至少要有16个0
const targetBit = 18 //挖矿难度

//ProofOfWork 工作量证明对象
type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数据存储，区块难度
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(pow.Block.Height),
		},
		[]byte{})

	return data
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1. 讲Block的属性拼接成字节数组

	//2. 生成hash
	nonce := 0
	//3. 循环判断hash有效性，如果满足条件，跳出循环

	var hashInt = new(big.Int) //用于存储新生成的hash
	var hash [32]byte
	for {
		//准备数据
		dataBytes := pow.prepareData(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		//讲hash存储到hashInt
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(hashInt) == 1 { //挖矿成功
			println(fmt.Sprintf("\r%v 挖矿成功：%x\n", nonce, hash))
			break
		}
		nonce++
	}
	return hash[:], int64(nonce)
}

func NewProofOfWork(block *Block) *ProofOfWork {

	//1. 创建一个初始值位1的target
	target := big.NewInt(1)
	//2. 左移256位 - targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{Block: block, target: target}
}
