package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

//Block 区块
type Block struct {
	Height        int64  //区块高度
	PrevBlockHash []byte //上一区块Hash
	Data          []byte //交易数据
	Timestamp     int64  //时间戳
	Hash          []byte //Hash
	Nonce         int64
}

// Serialize 序列化成字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeBlock 反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var res *Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}

	//	调用工作量证明返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	// 挖矿验证
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//CreateGenesisBlock 生成创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, bytes.Repeat([]byte{0}, 32))
}
