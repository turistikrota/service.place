package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/delivery/http/dto"
)

func (h Server) FeatureCreate(ctx *fiber.Ctx) error {
	d := dto.Request.FeatureCreate()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.FeatureCreate.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.FeatureCreate)
}

func (h Server) FeatureUpdate(ctx *fiber.Ctx) error {
	detail := dto.Request.FeatureDetail()
	d := dto.Request.FeatureUpdate()
	h.parseParams(ctx, detail)
	d.Load(detail)
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

func (h Server) AdminFeatureList(ctx *fiber.Ctx) error {
	res, err := h.app.Queries.AdminFeatureListAll.Handle(ctx.UserContext(), query.AdminFeatureListAllQuery{})
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.FeatureList, func() interface{} {
		return dto.Response.AdminFeatureList(res)
	})
}

func (h Server) AdminFeatureDetail(ctx *fiber.Ctx) error {
	detail := dto.Request.FeatureDetail()
	h.parseParams(ctx, detail)
	res, err := h.app.Queries.AdminFeatureDetail.Handle(ctx.UserContext(), detail.ToQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.FeatureList, func() interface{} {
		return dto.Response.AdminFeatureDetail(res)
	})
}

func (h Server) PlaceCreate(ctx *fiber.Ctx) error {
	d := dto.Request.PlaceCreate()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlaceCreate.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlaceCreate)
}

func (h Server) PlaceUpdate(ctx *fiber.Ctx) error {
	detail := dto.Request.PlaceDetail()
	d := dto.Request.PlaceUpdate()
	h.parseParams(ctx, detail)
	d.Load(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlaceUpdate.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlaceUpdate)
}

func (h Server) PlaceDelete(ctx *fiber.Ctx) error {
	detail := dto.Request.PlaceDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.PlaceDelete.Handle(ctx.UserContext(), detail.ToDeleteCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlaceDelete)
}

func (h Server) PlaceDisable(ctx *fiber.Ctx) error {
	detail := dto.Request.PlaceDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.PlaceDisable.Handle(ctx.UserContext(), detail.ToDisableCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlaceDisable)
}

func (h Server) PlaceEnable(ctx *fiber.Ctx) error {
	detail := dto.Request.PlaceDetail()
	h.parseParams(ctx, detail)
	_, err := h.app.Commands.PlaceEnable.Handle(ctx.UserContext(), detail.ToEnableCommand())
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlaceEnable)
}

func (h Server) PlaceFilter(ctx *fiber.Ctx) error {
	p := dto.Request.Pagination()
	d := dto.Request.PlaceFilter()
	h.parseQuery(ctx, p)
	d.LoadPagination(p)
	h.parseBody(ctx, d)
	l, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	res, err := h.app.Queries.PlaceFilter.Handle(ctx.UserContext(), d.ToQuery(l))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlaceFilter, func() interface{} {
		return dto.Response.PlaceList(res)
	})
}

func (h Server) PlaceAdminFilter(ctx *fiber.Ctx) error {
	p := dto.Request.Pagination()
	d := dto.Request.PlaceFilter()
	h.parseQuery(ctx, p)
	d.LoadPagination(p)
	h.parseBody(ctx, d)
	l, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	res, err := h.app.Queries.PlaceAdminFilter.Handle(ctx.UserContext(), d.ToAdminQuery(l))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlaceFilter, func() interface{} {
		return dto.Response.PlaceAdminList(res)
	})
}

func (h Server) PlaceView(ctx *fiber.Ctx) error {
	p := dto.Request.PlaceView()
	h.parseParams(ctx, p)
	l, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	res, err := h.app.Queries.PlaceView.Handle(ctx.UserContext(), p.ToQuery(l))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlaceView, func() interface{} {
		return dto.Response.PlaceView(res)
	})
}

func (h Server) PlaceAdminView(ctx *fiber.Ctx) error {
	detail := dto.Request.PlaceDetail()
	h.parseParams(ctx, detail)
	res, err := h.app.Queries.PlaceAdminView.Handle(ctx.UserContext(), detail.ToQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlaceView, func() interface{} {
		return dto.Response.PlaceAdminView(res)
	})
}
