package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mixarchitecture/microp/server/http"
	"github.com/mixarchitecture/microp/server/http/parser"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.place/src/app"
	"github.com/turistikrota/service.place/src/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
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

	feature := router.Group("/feature")

	feature.Get("/", h.wrapWithTimeout(h.FeatureList))
	featureAdmin := feature.Group("/", h.currentUserAccess(), h.requiredAccess())
	featureAdmin.Post("/", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureCreate), h.wrapWithTimeout(h.FeatureCreate))
	featureAdmin.Put("/:uuid/disable", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureDisable), h.wrapWithTimeout(h.FeatureDisable))
	featureAdmin.Put("/:uuid/enable", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureEnable), h.wrapWithTimeout(h.FeatureEnable))
	featureAdmin.Put("/:uuid", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureUpdate), h.wrapWithTimeout(h.FeatureUpdate))
	featureAdmin.Delete("/:uuid", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureDelete), h.wrapWithTimeout(h.FeatureDelete))
	featureAdmin.Get("/all", h.adminRoute(config.Roles.FeatureAll, config.Roles.FeatureList), h.wrapWithTimeout(h.AdminFeatureList))

	placeAdmin := router.Group("/place", h.currentUserAccess(), h.requiredAccess())
	placeAdmin.Post("/", h.adminRoute(config.Roles.PlaceAll, config.Roles.PlaceCreate), h.wrapWithTimeout(h.PlaceCreate))
	placeAdmin.Put("/:uuid/disable", h.adminRoute(config.Roles.PlaceAll, config.Roles.PlaceDisable), h.wrapWithTimeout(h.PlaceDisable))
	placeAdmin.Put("/:uuid/enable", h.adminRoute(config.Roles.PlaceAll, config.Roles.PlaceEnable), h.wrapWithTimeout(h.PlaceEnable))
	placeAdmin.Put("/:uuid", h.adminRoute(config.Roles.PlaceAll, config.Roles.PlaceUpdate), h.wrapWithTimeout(h.PlaceUpdate))
	placeAdmin.Delete("/:uuid", h.adminRoute(config.Roles.PlaceAll, config.Roles.PlaceDelete), h.wrapWithTimeout(h.PlaceDelete))

	router.Post("/", h.wrapWithTimeout(h.PlaceFilter))
	router.Get("/:slug", h.wrapWithTimeout(h.PlaceView))

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
