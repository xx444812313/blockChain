package main

import (
	"blockChain/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	blockchain.AddBlockToBlockchain("Weng Xiao Song 100RMB", 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Weng Xiao Song 2020RMB", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Weng Xiao Song 400RMB", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Weng Xiao Song 800RMB", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Weng Xiao Song 3992RMB", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Weng Xiao Song 20050RMB", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)


}
