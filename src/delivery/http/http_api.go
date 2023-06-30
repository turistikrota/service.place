package http

import (
	"api.turistikrota.com/place/src/delivery/http/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.shared/server/http/result"
)

func (h Server) FeatureCreate(ctx *fiber.Ctx) error {
	d := dto.Request.FeatureCreate()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.FeatureCreate.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureCreate)
}

func (h Server) FeatureUpdate(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) FeatureDelete(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) FeatureDisable(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) FeatureEnable(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) FeatureList(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceCreate(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceUpdate(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceDelete(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceDisable(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceEnable(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceFilter(ctx *fiber.Ctx) error {
	return nil
}

func (h Server) PlaceView(ctx *fiber.Ctx) error {
	return nil
}
