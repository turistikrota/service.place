package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	PlatformViewDetail        string
	PlatformCreate            string
	PlatformUpdate            string
	PlatformDelete            string
	PlatformEnable            string
	PlatformDisable           string
	PlatformTranslationAdd    string
	PlatformTranslationUpdate string
	PlatformTranslationDelete string
}

var Roles = roles{
	Roles:                     base_roles.BaseRoles,
	PlatformViewDetail:        "platform.view.detail",
	PlatformCreate:            "platform.create",
	PlatformUpdate:            "platform.update",
	PlatformDelete:            "platform.delete",
	PlatformEnable:            "platform.enable",
	PlatformDisable:           "platform.disable",
	PlatformTranslationAdd:    "platform.translation.add",
	PlatformTranslationUpdate: "platform.translation.update",
	PlatformTranslationDelete: "platform.translation.delete",
}
