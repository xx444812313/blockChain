package main

import (
	"blockChain/BLC"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	chain := BLC.CreateBlockchainWithGenesisBlock()
	spew.Dump(chain)

}
