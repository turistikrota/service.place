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
	PlaceAll       string
	PlaceCreate    string
	PlaceUpdate    string
	PlaceDelete    string
	PlaceEnable    string
	PlaceDisable   string
}

var Roles = roles{
	Roles:          base_roles.BaseRoles,
	FeatureAll:     "place.feature.all",
	FeatureCreate:  "place.feature.create",
	FeatureUpdate:  "place.feature.update",
	FeatureDelete:  "place.feature.delete",
	FeatureEnable:  "place.feature.enable",
	FeatureDisable: "place.feature.disable",
	PlaceAll:       "place.all",
	PlaceCreate:    "place.create",
	PlaceUpdate:    "place.update",
	PlaceDelete:    "place.delete",
	PlaceEnable:    "place.enable",
	PlaceDisable:   "place.disable",
}
