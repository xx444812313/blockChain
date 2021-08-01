package BLC

import "math/big"

//ProofOfWork 工作量证明对象
type ProofOfWork struct {
	Block  *Block //当前要验证的区块
	target *big.Int
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}

func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{Block: block}
}
