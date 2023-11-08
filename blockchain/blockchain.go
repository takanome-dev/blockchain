package blockchain

type BlockChain struct {
    Chain           []Block
    PendingTransactions []Transaction
}

func (b *BlockChain) GetLastBlock() Block {
    return b.Chain[len(b.Chain)-1]
}

func (b *BlockChain) CreateNewTransaction(amount int, sender, recipient string) int {
    tx := NewTransaction(amount, sender, recipient)
    b.PendingTransactions = append(b.PendingTransactions, *tx)
    return b.GetLastBlock().Index + 1
}












