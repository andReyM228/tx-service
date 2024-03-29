package delivery

import (
	"tx_service/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func Respond(ctx *fiber.Ctx, statusCode int, payload interface{}) error {
	ctx.Response().SetStatusCode(statusCode)

	if err := ctx.JSON(payload); err != nil {
		return err
	}

	return nil
}

func HandleError(ctx *fiber.Ctx, err error) error {
	switch err.(type) {
	case repositories.NotFound:
		if err := Respond(ctx, fiber.StatusNotFound, err); err != nil {
			return err
		}
	default:
		if err := Respond(ctx, fiber.StatusInternalServerError, err); err != nil {
			return err
		}

	}

	return nil
}
