package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

func (chain *Blockchain) AddBlockToBlockchain(data string, height int64, prevHash []byte) {
	newBlock := NewBlock(data, height, prevHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func CreateBlockchainWithGenesisBlock() *Blockchain {
	blockchain := &Blockchain{}
	genesisBlock := CreateGenesisBlock("Genesis Data......")
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
	return blockchain
}
