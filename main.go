package main

import (
	"blockChain/BLC"
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func main() {
	log.Println("hello blockChain")

	block := BLC.NewBlock("小松", 1, bytes.Repeat([]byte{0}, 32))
	log.Println(BLC.IntToHex(1))
	spew.Dump(block)
}
