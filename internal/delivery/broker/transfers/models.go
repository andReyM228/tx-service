package transfers

import (
	"github.com/andReyM228/lib/bus"
	"tx_service/internal/domain"
)

func toDomainIssue(transfer bus.IssueRequest) domain.Transaction {
	return domain.Transaction{
		ToAddress: transfer.ToAddress,
		Amount:    transfer.Amount,
		Memo:      transfer.Memo,
	}
}

func toDomainWithdraw(transfer bus.WithdrawRequest) domain.Transaction {
	return domain.Transaction{
		ToAddress: transfer.ToAddress,
		Amount:    transfer.Amount,
		Memo:      transfer.Memo,
	}

}

func toTransferResponse(txHash string) bus.TxResponse {
	return bus.TxResponse{
		TxHash: txHash,
	}
}
