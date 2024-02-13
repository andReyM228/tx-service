package handler

import "github.com/gofiber/fiber/v2"

type (
	Balances interface {
		Issue(ctx *fiber.Ctx) error
		Withdraw(ctx *fiber.Ctx) error
	}
)
