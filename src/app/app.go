package app

import (
	"api.turistikrota.com/place/src/app/command"
	"api.turistikrota.com/place/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AccountDelete             command.AccountDeleteHandler
	AccountCreate             command.AccountCreateHandler
	AccountUpdate             command.AccountUpdateHandler
	AccountEnable             command.AccountEnableHandler
	AccountDisable            command.AccountDisableHandler
	AccountSocialAdd          command.AccountSocialAddHandler
	AccountSocialRemove       command.AccountSocialRemoveHandler
	AccountSocialUpdate       command.AccountSocialUpdateHandler
	PlatformCreate            command.PlatformCreateHandler
	PlatformUpdate            command.PlatformUpdateHandler
	PlatformDisable           command.PlatformDisableHandler
	PlatformEnable            command.PlatformEnableHandler
	PlatformDelete            command.PlatformDeleteHandler
	PlatformTranslationCreate command.PlatformTranslationCreateHandler
	PlatformTranslationUpdate command.PlatformTranslationUpdateHandler
	PlatformTranslationRemove command.PlatformTranslationRemoveHandler
}

type Queries struct {
	AccountGet         query.AccountGetHandler
	AccountProfileView query.AccountProfileViewHandler
	AccountListMy      query.AccountListMyHandler
	PlatformGetBySlug  query.PlatformGetBySlugHandler
	PlatformListAll    query.PlatformListAllHandler
}
