package service

import (
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.place/src/adapters"
	"github.com/turistikrota/service.place/src/app"
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/config"
	"github.com/turistikrota/service.place/src/domain/feature"
	"github.com/turistikrota/service.place/src/domain/place"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {
	placeFactory := place.NewFactory()
	placeRepo := adapters.Mongo.NewPlace(placeFactory, cnf.Mongo.GetCollection(cnf.App.DB.Place.Collection))
	placeEvents := place.NewEvents(place.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
	})

	featureFactory := feature.NewFactory()
	featureRepo := adapters.Mongo.NewFeature(featureFactory, cnf.Mongo.GetCollection(cnf.App.DB.Feature.Collection))

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			FeatureCreate: command.NewFeatureCreateHandler(command.FeatureCreateHandlerConfig{
				Repo:     featureRepo,
				Factory:  featureFactory,
				CqrsBase: base,
			}),
			FeatureUpdate: command.NewFeatureUpdateHandler(command.FeatureUpdateHandlerConfig{
				Repo:     featureRepo,
				Factory:  featureFactory,
				CqrsBase: base,
			}),
			FeatureDelete: command.NewFeatureDeleteHandler(command.FeatureDeleteHandlerConfig{
				Repo:     featureRepo,
				Factory:  featureFactory,
				CqrsBase: base,
			}),
			FeatureDisable: command.NewFeatureDisableHandler(command.FeatureDisableHandlerConfig{
				Repo:     featureRepo,
				Factory:  featureFactory,
				CqrsBase: base,
			}),
			FeatureEnable: command.NewFeatureEnableHandler(command.FeatureEnableHandlerConfig{
				Repo:     featureRepo,
				Factory:  featureFactory,
				CqrsBase: base,
			}),
			PlaceCreate: command.NewPlaceCreateHandler(command.PlaceCreateHandlerConfig{
				Repo:        placeRepo,
				FeatureRepo: featureRepo,
				Events:      placeEvents,
				Factory:     placeFactory,
				CqrsBase:    base,
			}),
			PlaceUpdate: command.NewPlaceUpdateHandler(command.PlaceUpdateHandlerConfig{
				Repo:        placeRepo,
				Events:      placeEvents,
				FeatureRepo: featureRepo,
				Factory:     placeFactory,
				CqrsBase:    base,
			}),
			PlaceDelete: command.NewPlaceDeleteHandler(command.PlaceDeleteHandlerConfig{
				Repo:     placeRepo,
				Events:   placeEvents,
				Factory:  placeFactory,
				CqrsBase: base,
			}),
			PlaceDisable: command.NewPlaceDisableHandler(command.PlaceDisableHandlerConfig{
				Repo:     placeRepo,
				Events:   placeEvents,
				Factory:  placeFactory,
				CqrsBase: base,
			}),
			PlaceEnable: command.NewPlaceEnableHandler(command.PlaceEnableHandlerConfig{
				Repo:     placeRepo,
				Events:   placeEvents,
				Factory:  placeFactory,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{
			FeatureListAll: query.NewFeatureListAllHandler(query.FeatureListAllHandlerConfig{
				Repo:     featureRepo,
				CacheSrv: cnf.CacheSrv,
				CqrsBase: base,
			}),
			AdminFeatureListAll: query.NewAdminFeatureListAllHandler(query.AdminFeatureListAllHandlerConfig{
				Repo:     featureRepo,
				CqrsBase: base,
			}),
			AdminFeatureDetail: query.NewAdminFeatureDetailHandler(query.AdminFeatureDetailHandlerConfig{
				Repo:     featureRepo,
				CqrsBase: base,
			}),
			PlaceFilter: query.NewPlaceFilterHandler(query.PlaceFilterHandlerConfig{
				Repo:     placeRepo,
				CacheSrv: cnf.CacheSrv,
				CqrsBase: base,
			}),
			PlaceAdminFilter: query.NewPlaceAdminFilterHandler(query.PlaceAdminFilterHandlerConfig{
				Repo:     placeRepo,
				CacheSrv: cnf.CacheSrv,
				CqrsBase: base,
			}),
			PlaceView: query.NewPlaceViewHandler(query.PlaceViewHandlerConfig{
				Repo:        placeRepo,
				FeatureRepo: featureRepo,
				CacheSrv:    cnf.CacheSrv,
				CqrsBase:    base,
			}),
		},
	}
}
