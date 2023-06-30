package http

import (
	"api.turistikrota.com/place/src/app/query"
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
	d := dto.Request.FeatureUpdate()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.FeatureUpdate.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureUpdate)
}

func (h Server) FeatureDelete(ctx *fiber.Ctx) error {
	detail := dto.Request.FeatureDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.FeatureDelete.Handle(ctx.UserContext(), detail.ToDeleteCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureDelete)
}

func (h Server) FeatureDisable(ctx *fiber.Ctx) error {
	detail := dto.Request.FeatureDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.FeatureDisable.Handle(ctx.UserContext(), detail.ToDisableCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureDisable)
}

func (h Server) FeatureEnable(ctx *fiber.Ctx) error {
	detail := dto.Request.FeatureDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.FeatureEnable.Handle(ctx.UserContext(), detail.ToEnableCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureEnable)
}

func (h Server) FeatureList(ctx *fiber.Ctx) error {
	res, err := h.app.Queries.FeatureListAll.Handle(ctx.UserContext(), query.FeatureListAllQuery{})
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.FeatureList, func() interface{} {
		return dto.Response.FeatureList(res)
	})
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
