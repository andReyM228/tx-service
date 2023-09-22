package balances

import (
	"tx_service/internal/domain"
	"tx_service/internal/handler"
	balancesService "tx_service/internal/service/balances"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	balances balancesService.Service
}

func NewHandler(balances balancesService.Service) Handler {
	return Handler{
		balances: balances,
	}
}

func (h Handler) Transfer(ctx *fiber.Ctx) error {
	var tx domain.Transactions
	if err := ctx.BodyParser(&tx); err != nil {
		return handler.HandleError(ctx, err)
	}

	if err := h.balances.SendCoins(tx); err != nil {
		return handler.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
