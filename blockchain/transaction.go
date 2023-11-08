package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Transaction struct {
	ID        []byte
	Amount    int
	Sender    string
	Recipient string
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	HandleErr(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func NewTransaction(amount int, sender, recipient string) *Transaction {
	tx := Transaction{
		ID: []byte{},
		Amount: amount,
		Sender:	sender,
		Recipient: recipient,
	}
	tx.SetID()

	return &tx
}