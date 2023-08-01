package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	FeatureAll     string
	FeatureCreate  string
	FeatureUpdate  string
	FeatureDelete  string
	FeatureEnable  string
	FeatureDisable string
	FeatureList    string
	FeatureRead    string
	PlaceAll       string
	PlaceCreate    string
	PlaceList      string
	PlaceUpdate    string
	PlaceDelete    string
	PlaceEnable    string
	PlaceDisable   string
	PlaceView      string
}

var Roles = roles{
	Roles:          base_roles.BaseRoles,
	FeatureAll:     "place.feature.all",
	FeatureCreate:  "place.feature.create",
	FeatureUpdate:  "place.feature.update",
	FeatureDelete:  "place.feature.delete",
	FeatureEnable:  "place.feature.enable",
	FeatureDisable: "place.feature.disable",
	FeatureList:    "place.feature.list",
	FeatureRead:    "place.feature.read",
	PlaceAll:       "place.all",
	PlaceList:      "place.list",
	PlaceCreate:    "place.create",
	PlaceUpdate:    "place.update",
	PlaceDelete:    "place.delete",
	PlaceEnable:    "place.enable",
	PlaceDisable:   "place.disable",
	PlaceView:      "place.view",
}
