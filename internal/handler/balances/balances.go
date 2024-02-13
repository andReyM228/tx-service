package balances

import (
	"github.com/gofiber/fiber/v2"
	"tx_service/internal/handler"
	"tx_service/internal/service"
)

type Handler struct {
	balances service.Balances
}

func NewHandler(balances service.Balances) Handler {
	return Handler{
		balances: balances,
	}
}

func (h Handler) Issue(ctx *fiber.Ctx) error {
	var request transaction

	if err := ctx.BodyParser(&request); err != nil {
		return handler.HandleError(ctx, err)
	}

	txHash, err := h.balances.Issue(ctx.Context(), toDomain(request))
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	return handler.Respond(ctx, fiber.StatusOK, transactionHash{TxHash: txHash})
}

func (h Handler) Withdraw(ctx *fiber.Ctx) error {
	var request transaction

	if err := ctx.BodyParser(&request); err != nil {
		return handler.HandleError(ctx, err)
	}

	txHash, err := h.balances.Withdraw(ctx.Context(), toDomain(request))
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	return handler.Respond(ctx, fiber.StatusOK, transactionHash{TxHash: txHash})
}
