package delivery

import (
	"context"

	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"

	"api.turistikrota.com/account/src/app"
	"api.turistikrota.com/account/src/config"
	"api.turistikrota.com/account/src/delivery/event_stream"
	"api.turistikrota.com/account/src/delivery/http"
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/events"
	sharedHttp "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/validator"
)

type Delivery interface {
	Load()
}

type delivery struct {
	app         app.Application
	config      config.App
	i18n        *i18np.I18n
	ctx         context.Context
	validator   *validator.Validator
	eventEngine events.Engine
	sessionSrv  session.Service
	tknSrv      token.Service
}

type Config struct {
	App         app.Application
	Config      config.App
	I18n        *i18np.I18n
	Validator   *validator.Validator
	Ctx         context.Context
	EventEngine events.Engine
	SessionSrv  session.Service
	TokenSrv    token.Service
}

func New(config Config) Delivery {
	return &delivery{
		app:         config.App,
		config:      config.Config,
		i18n:        config.I18n,
		ctx:         config.Ctx,
		eventEngine: config.EventEngine,
		sessionSrv:  config.SessionSrv,
		tknSrv:      config.TokenSrv,
		validator:   config.Validator,
	}
}

func (d *delivery) Load() {
	d.loadEventStream().loadHTTP()
}

func (d *delivery) loadHTTP() *delivery {
	sharedHttp.RunServer(sharedHttp.Config{
		Host:  d.config.Server.Host,
		Port:  d.config.Server.Port,
		I18n:  d.i18n,
		Group: d.config.Server.Group,
		CreateHandler: func(router fiber.Router) fiber.Router {
			return http.New(http.Config{
				App:         d.app,
				I18n:        *d.i18n,
				Validator:   *d.validator,
				Context:     d.ctx,
				HttpHeaders: d.config.HttpHeaders,
				SessionSrv:  d.sessionSrv,
				TokenSrv:    d.tknSrv,
			}).Load(router)
		},
	})
	return d
}

func (d *delivery) loadEventStream() *delivery {
	eventStream := event_stream.New(event_stream.Config{
		App:    d.app,
		Engine: d.eventEngine,
	})
	err := d.eventEngine.Open()
	if err != nil {
		panic(err)
	}
	eventStream.Load()
	return d
}
