package http

import (
	"api.turistikrota.com/account/src/delivery/http/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/result"
)

func (h Server) AccountDelete(ctx *fiber.Ctx) error {
	d := dto.Request.AccountDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.AccountDelete.Handle(ctx.UserContext(), d.ToDeleteCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountDelete)
}

func (h Server) AccountCreate(ctx *fiber.Ctx) error {
	d := dto.Request.AccountCreate()
	h.parseBody(ctx, d)
	res, err := h.app.Commands.AccountCreate.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.AccountCreate, func() interface{} {
		return dto.Response.AccountCreate(res)
	})
}

func (h Server) AccountUpdate(ctx *fiber.Ctx) error {
	detail := dto.Request.AccountDetail()
	d := dto.Request.AccountUpdate()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.AccountUpdate.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountUpdate)
}

func (h Server) AccountDisable(ctx *fiber.Ctx) error {
	d := dto.Request.AccountDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.AccountDisable.Handle(ctx.UserContext(), d.ToDisableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountDisable)
}

func (h Server) AccountEnable(ctx *fiber.Ctx) error {
	d := dto.Request.AccountDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.AccountEnable.Handle(ctx.UserContext(), d.ToEnableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountEnable)
}

func (h Server) AccountSocialAdd(ctx *fiber.Ctx) error {
	social := dto.Request.AccountSocial()
	d := dto.Request.AccountSocialAction()
	h.parseParams(ctx, social)
	d.LoadSocial(social)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.AccountSocialAdd.Handle(ctx.UserContext(), d.ToAddCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountSocialAdd)
}

func (h Server) AccountSocialUpdate(ctx *fiber.Ctx) error {
	social := dto.Request.AccountSocial()
	d := dto.Request.AccountSocialAction()
	h.parseParams(ctx, social)
	d.LoadSocial(social)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.AccountSocialUpdate.Handle(ctx.UserContext(), d.ToUpdateCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountSocialUpdate)
}

func (h Server) AccountSocialRemove(ctx *fiber.Ctx) error {
	d := dto.Request.AccountSocial()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.AccountSocialRemove.Handle(ctx.UserContext(), d.ToRemoveCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.AccountSocialRemove)
}

func (h Server) AccountGet(ctx *fiber.Ctx) error {
	d := dto.Request.AccountDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.AccountGet.Handle(ctx.UserContext(), d.ToGetQuery(current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.AccountGet, func() interface{} {
		return dto.Response.AccountGet(res)
	})
}

func (h Server) AccountListMy(ctx *fiber.Ctx) error {
	d := dto.Request.AccountListMy()
	res, err := h.app.Queries.AccountListMy.Handle(ctx.UserContext(), d.ToQuery(current_user.Parse(ctx).UUID))
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.AccountListMy, func() interface{} {
		return dto.Response.AccountListMy(res)
	})
}

func (h Server) AccountProfileView(ctx *fiber.Ctx) error {
	d := dto.Request.AccountDetail()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.AccountProfileView.Handle(ctx.UserContext(), d.ToProfileQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.AccountProfileView, func() interface{} {
		return dto.Response.AccountProfileView(res)
	})
}

func (h Server) PlatformCreate(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformCreate()
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlatformCreate.Handle(ctx.UserContext(), d.ToCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformCreate)
}

func (h Server) PlatformUpdate(ctx *fiber.Ctx) error {
	detail := dto.Request.PlatformDetail()
	d := dto.Request.PlatformAction()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlatformUpdate.Handle(ctx.UserContext(), d.ToUpdateCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformUpdate)
}

func (h Server) PlatformDelete(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.PlatformDelete.Handle(ctx.UserContext(), d.ToDeleteCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformDelete)
}

func (h Server) PlatformDisable(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.PlatformDisable.Handle(ctx.UserContext(), d.ToDisableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformDisable)
}

func (h Server) PlatformEnable(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformDetail()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.PlatformEnable.Handle(ctx.UserContext(), d.ToEnableCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformEnable)
}

func (h Server) PlatformGet(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformGet()
	h.parseParams(ctx, d)
	res, err := h.app.Queries.PlatformGetBySlug.Handle(ctx.UserContext(), d.ToQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlatformGet, func() interface{} {
		return dto.Response.PlatformGet(res)
	})
}

func (h Server) PlatformList(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformList()
	res, err := h.app.Queries.PlatformListAll.Handle(ctx.UserContext(), d.ToQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.PlatformList, func() interface{} {
		return dto.Response.PlatformList(res)
	})
}

func (h Server) PlatformTranslationCreate(ctx *fiber.Ctx) error {
	detail := dto.Request.PlatformTranslation()
	d := dto.Request.PlatformTranslationAction()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlatformTranslationCreate.Handle(ctx.UserContext(), d.ToCreateCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformTranslationCreate)
}

func (h Server) PlatformTranslationUpdate(ctx *fiber.Ctx) error {
	detail := dto.Request.PlatformTranslation()
	d := dto.Request.PlatformTranslationAction()
	h.parseParams(ctx, detail)
	d.LoadDetail(detail)
	h.parseBody(ctx, d)
	_, err := h.app.Commands.PlatformTranslationUpdate.Handle(ctx.UserContext(), d.ToUpdateCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformTranslationUpdate)
}

func (h Server) PlatformTranslationDelete(ctx *fiber.Ctx) error {
	d := dto.Request.PlatformTranslation()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.PlatformTranslationRemove.Handle(ctx.UserContext(), d.ToRemoveCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.PlatformTranslationDelete)
}
