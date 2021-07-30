package main

import (
	"blockChain/BLC"
	"bytes"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	block := BLC.NewBlock("小松", 1, bytes.Repeat([]byte{0}, 32))
	spew.Dump(block)

	genesisBlock := BLC.CreateGenesisBlock("小松")
	spew.Dump(genesisBlock)

}
