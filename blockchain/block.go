package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []*Transaction
	Nonce        int
	PrevHash     []byte
	Hash         []byte
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
			txHashes = append(txHashes, []byte(tx.ID))
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func CreateBlock(tx []*Transaction, prevHash []byte, chainLength int) *Block {
	newBlock := Block{
			Index: chainLength + 1,
			Timestamp: time.Now(),
			Transactions: tx,
			Nonce: 0,
			PrevHash: prevHash,
			Hash: []byte{},
	}

	pow := NewProof(&newBlock)

	nonce, hash := pow.Run()

	newBlock.Hash = hash
	newBlock.Nonce = nonce

	return &newBlock
} 


func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}