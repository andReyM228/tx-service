package chain

import (
	"context"
	"github.com/andReyM228/lib/log"
	"github.com/andReyM228/one/chain_client"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type Repository struct {
	chain chain_client.Client
	log   log.Logger
}

func NewRepository(chain chain_client.Client, log log.Logger) Repository {
	return Repository{
		chain: chain,
		log:   log,
	}
}

func (r Repository) Send(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error) {
	txResp, err := r.chain.Send(ctx, toAddress, amount, memo, chain_client.DenomOne, signBy)
	if err != nil {
		return "", err
	}

	return txResp.TxHash, nil
}

func (r Repository) Withdraw(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error) {
	txResp, err := r.chain.Withdraw(ctx, toAddress, amount, memo, chain_client.DenomOne, signBy)
	if err != nil {
		return "", err
	}

	return txResp.TxHash, nil
}

func (r Repository) Issue(ctx context.Context, toAddress string, amount int64, memo string, signBy string) (string, error) {
	txResp, err := r.chain.Issue(ctx, toAddress, amount, memo, chain_client.DenomOne, signBy)
	if err != nil {
		return "", err
	}

	return txResp.TxHash, nil
}

func (r Repository) GetTxByHash(ctx context.Context, hash string) (*tx.GetTxResponse, error) {
	txResp, err := r.chain.GetTx(ctx, hash)
	if err != nil {
		return nil, err
	}

	return txResp, nil
}
