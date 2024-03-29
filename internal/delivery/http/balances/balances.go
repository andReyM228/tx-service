package balances

import (
	"github.com/gofiber/fiber/v2"
	"tx_service/internal/delivery"
	"tx_service/internal/services"
)

type Handler struct {
	balances services.Balances
}

func NewHandler(balances services.Balances) Handler {
	return Handler{
		balances: balances,
	}
}

func (h Handler) Issue(ctx *fiber.Ctx) error {
	var request transaction

	if err := ctx.BodyParser(&request); err != nil {
		return delivery.HandleError(ctx, err)
	}

	txHash, err := h.balances.Issue(ctx.Context(), toDomain(request))
	if err != nil {
		return delivery.HandleError(ctx, err)
	}

	return delivery.Respond(ctx, fiber.StatusOK, transactionHash{TxHash: txHash})
}

func (h Handler) Withdraw(ctx *fiber.Ctx) error {
	var request transaction

	if err := ctx.BodyParser(&request); err != nil {
		return delivery.HandleError(ctx, err)
	}

	txHash, err := h.balances.Withdraw(ctx.Context(), toDomain(request))
	if err != nil {
		return delivery.HandleError(ctx, err)
	}

	return delivery.Respond(ctx, fiber.StatusOK, transactionHash{TxHash: txHash})
}
