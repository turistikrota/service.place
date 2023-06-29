package service

import (
	"api.turistikrota.com/place/src/app"
	"api.turistikrota.com/place/src/config"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"
	"github.com/turistikrota/service.shared/decorator"
	"github.com/turistikrota/service.shared/events"
	"github.com/turistikrota/service.shared/validator"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
	CacheSrv    redis.Service
}

func NewApplication(config Config) app.Application {
	_ = decorator.NewBase()

	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
