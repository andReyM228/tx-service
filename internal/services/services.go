package services

import (
	"context"
	"tx_service/internal/domain"
)

type (
	Balances interface {
		Issue(ctx context.Context, req domain.Transaction) (string, error)
		Withdraw(ctx context.Context, req domain.Transaction) (string, error)
	}
)
