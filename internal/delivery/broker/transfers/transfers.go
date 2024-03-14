package transfers

import (
	"context"
	"encoding/json"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/log"
	"github.com/andReyM228/lib/rabbit"
	"tx_service/internal/delivery"
	"tx_service/internal/services"
)

type handler struct {
	rabbit  rabbit.Rabbit
	log     log.Logger
	service services.Balances
}

func NewHandler(rabbit rabbit.Rabbit, log log.Logger, service services.Balances) delivery.TransfersBroker {
	return handler{
		rabbit:  rabbit,
		log:     log,
		service: service,
	}
}

func (h handler) BrokerIssue(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var issueTransfer bus.IssueRequest
	if err := json.Unmarshal(req.Payload, &issueTransfer); err != nil {
		return err
	}

	txHash, err := h.service.Issue(context.Background(), toDomainIssue(issueTransfer))
	if err != nil {
		return err
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, toTransferResponse(txHash))
}

func (h handler) BrokerWithdraw(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var withdrawTransfer bus.WithdrawRequest
	if err := json.Unmarshal(req.Payload, &withdrawTransfer); err != nil {
		return err
	}

	txHash, err := h.service.Withdraw(context.Background(), toDomainWithdraw(withdrawTransfer))
	if err != nil {
		return err
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, toTransferResponse(txHash))
}
