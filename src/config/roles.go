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
	FeatureAll:     "feature.all",
	FeatureCreate:  "feature.create",
	FeatureUpdate:  "feature.update",
	FeatureDelete:  "feature.delete",
	FeatureEnable:  "feature.enable",
	FeatureDisable: "feature.disable",
	PlaceAll:       "place.all",
	PlaceCreate:    "place.create",
	PlaceUpdate:    "place.update",
	PlaceDelete:    "place.delete",
	PlaceEnable:    "place.enable",
	PlaceDisable:   "place.disable",
}
