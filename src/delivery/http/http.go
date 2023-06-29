package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
	"github.com/turistikrota/service.shared/server/http/parser"

	"api.turistikrota.com/place/src/app"
	"api.turistikrota.com/place/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/validator"
)

type Server struct {
	app         app.Application
	i18n        i18np.I18n
	validator   validator.Validator
	ctx         context.Context
	tknSrv      token.Service
	sessionSrv  session.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	App         app.Application
	I18n        i18np.I18n
	Validator   validator.Validator
	Context     context.Context
	HttpHeaders config.HttpHeaders
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) Server {
	return Server{
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		ctx:         config.Context,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Use(h.cors(), h.deviceUUID())

	platform := router.Group("/platform")
	platform.Get("/", h.wrapWithTimeout(h.PlatformList))
	platform.Post("/", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformCreate), h.wrapWithTimeout(h.PlatformCreate))
	detail := platform.Group("/:slug")
	detail.Post("/translation/:locale", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformTranslationAdd), h.wrapWithTimeout(h.PlatformTranslationCreate))
	detail.Put("/translation/:locale", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformTranslationUpdate), h.wrapWithTimeout(h.PlatformTranslationUpdate))
	detail.Delete("/translation/:locale", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformTranslationDelete), h.wrapWithTimeout(h.PlatformTranslationDelete))
	detail.Put("/enable", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformEnable), h.wrapWithTimeout(h.PlatformEnable))
	detail.Put("/disable", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformDisable), h.wrapWithTimeout(h.PlatformDisable))
	detail.Delete("/", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformDelete), h.wrapWithTimeout(h.PlatformDelete))
	detail.Put("/", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformUpdate), h.wrapWithTimeout(h.PlatformUpdate))
	detail.Get("/", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.PlatformViewDetail), h.wrapWithTimeout(h.PlatformGet))

	router.Post("/", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountCreate))
	router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountListMy))
	router.Get("/@:userName/my", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountGet))
	router.Get("/@:userName", h.wrapWithTimeout(h.AccountProfileView))
	router.Put("/@:userName/enable", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountEnable))
	router.Put("/@:userName/disable", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountDisable))
	router.Post("/@:userName/social/:platform", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountSocialAdd))
	router.Delete("/@:userName/social/:platform", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountSocialRemove))
	router.Put("/@:userName/social/:platform", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountSocialUpdate))
	router.Put("/@:userName", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountUpdate))
	router.Delete("/@:userName", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.AccountDelete))
	return router
}

func (h Server) parseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, h.i18n, d)
}

func (h Server) parseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, h.i18n, d)
}

func (h Server) parseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, h.i18n, d)
}

func (h Server) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}

func (h Server) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       &h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  http.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h Server) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h Server) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h Server) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h Server) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.httpHeaders.AllowedOrigins,
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
	})
}
