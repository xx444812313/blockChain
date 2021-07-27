package BLC

import (
	"bytes"
	"crypto/sha256"
	"log"
	"strconv"
	"time"
)

//Block 区块
type Block struct {
	Height        int64  //区块高度
	PrevBlockHash []byte //上一区块Hash
	Data          []byte //交易数据
	Timestamp     int64  //时间戳
	Hash          []byte //Hash
}

//SetHash 按照Block struct字段的顺序填充bytes计算出hash
func (b *Block) SetHash() {
	// 1.Height	>	[]byte
	heightBytes := IntToHex(b.Height)
	// 2.Timestamp	>	[]byte
	timeBytes := []byte(strconv.FormatInt(b.Timestamp, 2))

	joinBytes := bytes.Join([][]byte{heightBytes, b.PrevBlockHash, b.Data, timeBytes, b.Hash}, []byte{})
	log.Println(joinBytes)
	sum256 := sha256.Sum256(joinBytes)
	b.Hash = sum256[:]

}

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}
	block.SetHash()
	return block
}
