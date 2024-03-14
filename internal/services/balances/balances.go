package balances

import (
	"context"
	"github.com/andReyM228/lib/log"
	"tx_service/internal/domain"
	"tx_service/internal/repositories"
	"tx_service/internal/services"
)

type service struct {
	chainRepo repositories.Chain
	log       log.Logger
}

func NewService(chainRepo repositories.Chain, log log.Logger) services.Balances {
	return service{
		chainRepo: chainRepo,
		log:       log,
	}
}

func (s service) Issue(ctx context.Context, req domain.Transaction) (string, error) {
	txHash, err := s.chainRepo.Issue(ctx, req.ToAddress, req.Amount, req.Memo, domain.SignerAccount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s service) Withdraw(ctx context.Context, req domain.Transaction) (string, error) {
	txHash, err := s.chainRepo.Withdraw(ctx, req.ToAddress, req.Amount, req.Memo, domain.SignerAccount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
