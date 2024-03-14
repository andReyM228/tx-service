package repositories

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type (
	Chain interface {
		Send(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error)
		Withdraw(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error)
		Issue(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error)
		GetTxByHash(ctx context.Context, hash string) (*tx.GetTxResponse, error)
	}
)
