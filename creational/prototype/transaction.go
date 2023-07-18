package prototype

import "strconv"

type Transaction struct {
	Sender      string
	Recipient   string
	Amount      int
	Description string
}

func (transaction *Transaction) Clone() Prototype {
	transactionClone := *transaction
	transactionClone.Description = ".:Copy:. " + transaction.Description
	return &transactionClone
}
func (transaction *Transaction) String() string {
	return "Sender: " + transaction.Sender + "\n Recipient: " + transaction.Recipient + "\n Amount: " + strconv.Itoa(transaction.Amount) + "\n Description: " + transaction.Description + "\n "
}
