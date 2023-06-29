package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/account"
	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountSocialRemoveCommand struct {
	UserUUID    string
	AccountName string
	Platform    string
}

type AccountSocialRemoveResult struct{}

type AccountSocialRemoveHandler decorator.CommandHandler[AccountSocialRemoveCommand, *AccountSocialRemoveResult]

type accountSocialRemoveHandler struct {
	platformRepo    platform.Repository
	platformFactory platform.Factory
	accountRepo     account.Repository
	accountFactory  account.Factory
	events          account.Events
}

type AccountSocialRemoveHandlerConfig struct {
	PlatformRepo    platform.Repository
	PlatformFactory platform.Factory
	AccountRepo     account.Repository
	AccountFactory  account.Factory
	Events          account.Events
	CqrsBase        decorator.Base
}

func NewAccountSocialRemoveHandler(config AccountSocialRemoveHandlerConfig) AccountSocialRemoveHandler {
	return decorator.ApplyCommandDecorators[AccountSocialRemoveCommand, *AccountSocialRemoveResult](
		accountSocialRemoveHandler{
			platformRepo:    config.PlatformRepo,
			platformFactory: config.PlatformFactory,
			accountRepo:     config.AccountRepo,
			accountFactory:  config.AccountFactory,
			events:          config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountSocialRemoveHandler) Handle(ctx context.Context, command AccountSocialRemoveCommand) (*AccountSocialRemoveResult, *i18np.Error) {
	err := h.accountRepo.SocialRemove(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, command.Platform)
	if err != nil {
		return nil, err
	}
	h.events.SocialRemoved(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, command.Platform)
	return &AccountSocialRemoveResult{}, nil
}
