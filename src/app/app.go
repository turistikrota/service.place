package app

import (
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	FeatureCreate  command.FeatureCreateHandler
	FeatureUpdate  command.FeatureUpdateHandler
	FeatureDelete  command.FeatureDeleteHandler
	FeatureDisable command.FeatureDisableHandler
	FeatureEnable  command.FeatureEnableHandler
	PlaceCreate    command.PlaceCreateHandler
	PlaceUpdate    command.PlaceUpdateHandler
	PlaceDelete    command.PlaceDeleteHandler
	PlaceDisable   command.PlaceDisableHandler
	PlaceEnable    command.PlaceEnableHandler
}

type Queries struct {
	AdminFeatureListAll query.AdminFeatureListAllHandler
	AdminFeatureDetail  query.AdminFeatureDetailHandler
	FeatureListAll      query.FeatureListAllHandler
	PlaceFilter         query.PlaceFilterHandler
	PlaceAdminFilter    query.PlaceAdminFilterHandler
	PlaceView           query.PlaceViewHandler
}
