package blockchain

import (
	"fmt"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	tx := []*Transaction{
		{ Amount: 1000, Sender: "John", Recipient: "Anna"},
		{ Amount: 1500, Sender: "Brian", Recipient: "Brandon"},
		{ Amount: 2000, Sender: "Eva", Recipient: "Nick"},
	}

	block := CreateBlock(tx, []byte{}, 0)

	fmt.Println()
	fmt.Printf("⛓ new block created ⛓ -> %+v", block)
	fmt.Println()

	hashedTx := block.HashTransactions()
	fmt.Println()
	fmt.Printf("🔒 hashed transactions 🔒 -> %x", hashedTx)
	fmt.Println()
}