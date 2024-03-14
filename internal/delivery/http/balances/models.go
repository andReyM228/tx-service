package balances

import "tx_service/internal/domain"

type transaction struct {
	ToAddress string `json:"to_address"`
	Amount    int64  `json:"amount"`
	Memo      string `json:"memo"`
}

type transactionHash struct {
	TxHash string `json:"tx_hash"`
}

func toDomain(t transaction) domain.Transaction {
	return domain.Transaction{
		ToAddress: t.ToAddress,
		Amount:    t.Amount,
		Memo:      t.Memo,
	}
}
