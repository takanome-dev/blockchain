package blockchain

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"
)

func TestNewProof(t *testing.T) {
	transactions := []*Transaction{
		{ Amount: 1000, Sender: "John", Recipient: "Anna"},
		{ Amount: 1500, Sender: "Brian", Recipient: "Brandon"},
		{ Amount: 2000, Sender: "Eva", Recipient: "Nick"},
	}

	block := Block{
		Index: 1,
		Timestamp: time.Now(),
		Transactions: transactions,
		Nonce: 0,
		PrevHash: []byte{},
		Hash: sha256.New().Sum([]byte{}),
	}

	pow := NewProof(&block)
	fmt.Println()
	fmt.Printf("📜 proof of work created 📜 -> %+v\n", pow.Block)
	fmt.Println()

	nonce, hash := pow.Run()
	fmt.Println()
	fmt.Printf("📜 pow done 📜\nnonce -> %d\nhash -> %x\n", nonce, hash)
	fmt.Println()

	if nonce != 1305 {
		t.Fatalf("expected %d, got %d\n", 1035, nonce)
	}
}