package balances

import (
	"context"
	"github.com/andReyM228/lib/log"
	"tx_service/internal/domain"
	"tx_service/internal/repository"
)

type Service struct {
	chainRepo repository.Chain
	log       log.Logger
}

func NewService(chainRepo repository.Chain, log log.Logger) Service {
	return Service{
		chainRepo: chainRepo,
		log:       log,
	}
}

func (s Service) Issue(ctx context.Context, req domain.Transaction) (string, error) {
	txHash, err := s.chainRepo.Issue(ctx, req.ToAddress, req.Amount, req.Memo, domain.SignerAccount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s Service) Withdraw(ctx context.Context, req domain.Transaction) (string, error) {
	txHash, err := s.chainRepo.Withdraw(ctx, req.ToAddress, req.Amount, req.Memo, domain.SignerAccount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
