package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const (
	// 数据库名字
	dbName = "blockChain.db"
	// 表的名字
	blockTableName = "blocks"
	// 最新区块的Hash
	newestBlockHash = "newest_block_hash"
)

type Blockchain struct {
	Tip []byte   //最新区块的Hash
	DB  *bolt.DB //数据库对象
}

//func (chain *Blockchain) AddBlockToBlockchain(data string, height int64, prevHash []byte) {
//	newBlock := NewBlock(data, height, prevHash)
//	chain.Blocks = append(chain.Blocks, newBlock)
//}

func CreateBlockchainWithGenesisBlock() *Blockchain {
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}

	var blockHash []byte

	e = db.Update(func(tx *bolt.Tx) error {
		bucket, e := tx.CreateBucketIfNotExists([]byte(blockTableName))
		if e != nil {
			return e
		}
		//创建创世区块
		genesisBlock := CreateGenesisBlock("小松golang开发工程师")
		//存储创世区块
		e = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if e != nil {
			return e
		}
		//存储最新区块Hash
		blockHash = genesisBlock.Hash
		e = bucket.Put([]byte(newestBlockHash), blockHash)
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		log.Panic(e)
	}

	return &Blockchain{blockHash, db}
}
