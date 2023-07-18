package prototype

import "strconv"

type Block struct {
	TransactionList []Prototype
	PreviousHash    string
}

func (block *Block) Clone() Prototype {
	blockClone := *block
	transactionList := []Prototype{}
	for _, transaction := range block.TransactionList {
		transactionList = append(transactionList, transaction.Clone())
	}
	blockClone.TransactionList = transactionList
	return &blockClone
}
func (block *Block) String() string {
	transactionString := ""
	for _, transaction := range block.TransactionList {
		transactionString += transaction.String()
	}
	return "Transaction Count: " + strconv.Itoa(len(block.TransactionList)) + "\n PreviousHash: " + block.PreviousHash + "\n Transaction List:\n" + transactionString
}
