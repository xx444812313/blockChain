package db

import (
	"github.com/boltdb/bolt"
	"log"
)

func CreateDB(dbName string) (*bolt.DB, error) {
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		return nil, e
	}
	return db, nil
}

func CreateBucket() error {
	db, e := CreateDB("example.db")
	if e != nil {
		return e
	}
	defer db.Close()
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("blockChain"))
		//bucket, err := tx.CreateBucket([]byte("blockChain"))	//创建表

		err := bucket.Put([]byte("2"), []byte("Send 100 BTC To 小松......"))
		if err != nil {
			log.Panic("数据存储失败......")
		} else {
			log.Println("数据存储成功")
		}
		return nil
	})
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("blockChain")) //需要判断是否拿到表了，这里是不会给出err的
		getBytes := bucket.Get([]byte("2"))
		log.Println("读取到数据：", string(getBytes))
		e2 := bucket.Put([]byte("3"), []byte("小松2021年这一年要成为区块链开发高手..."))
		if e2 != nil {
			log.Println("写入失败，error=", e2)
		}

		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
	return err
}
