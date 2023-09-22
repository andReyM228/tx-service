package balances

import (
	"errors"

	"tx_service/internal/domain"
	"tx_service/internal/repository/balances"
	"tx_service/internal/repository/transactions"
)

type Service struct {
	balanceRepo     balances.Repository
	transactionRepo transactions.Repository
}

func NewService(balanceRepo balances.Repository, transactionRepo transactions.Repository) Service {
	return Service{
		balanceRepo:     balanceRepo,
		transactionRepo: transactionRepo,
	}
}

func (s Service) SendCoins(tx domain.Transactions) error {
	fromBalance, err := s.balanceRepo.Get(int64(tx.UserIDFrom))
	if err != nil {
		return err
	}

	if fromBalance.Amount == 0 {
		return errors.New("not enough balance")
	}

	//transactionCost, err := s.balanceRepo.Get(tx.Amount)
	//if err != nil {
	//	return err
	//}

	if fromBalance.Amount < tx.Amount {
		return errors.New("not enough balance for transaction")
	}

	toBalance, err := s.balanceRepo.Get(int64(tx.UserIDTo))
	if err != nil {
		return err
	}

	fromBalance.Amount = fromBalance.Amount - tx.Amount
	toBalance.Amount = toBalance.Amount + tx.Amount

	if err := s.balanceRepo.Update(fromBalance); err != nil {
		return err
	}

	if err = s.balanceRepo.Update(toBalance); err != nil {
		return err
	}

	if err = s.transactionRepo.Create(tx); err != nil {
		return err
	}

	return nil
}
