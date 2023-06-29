package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/account"
	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountSocialUpdateCommand struct {
	UserUUID    string
	AccountName string
	Platform    string
	Value       string
}

type AccountSocialUpdateResult struct{}

type AccountSocialUpdateHandler decorator.CommandHandler[AccountSocialUpdateCommand, *AccountSocialUpdateResult]

type accountSocialUpdateHandler struct {
	platformRepo    platform.Repository
	platformFactory platform.Factory
	accountRepo     account.Repository
	accountFactory  account.Factory
	events          account.Events
}

type AccountSocialUpdateHandlerConfig struct {
	PlatformRepo    platform.Repository
	PlatformFactory platform.Factory
	AccountRepo     account.Repository
	AccountFactory  account.Factory
	Events          account.Events
	CqrsBase        decorator.Base
}

func NewAccountSocialUpdateHandler(config AccountSocialUpdateHandlerConfig) AccountSocialUpdateHandler {
	return decorator.ApplyCommandDecorators[AccountSocialUpdateCommand, *AccountSocialUpdateResult](
		accountSocialUpdateHandler{
			platformRepo:    config.PlatformRepo,
			platformFactory: config.PlatformFactory,
			accountRepo:     config.AccountRepo,
			accountFactory:  config.AccountFactory,
			events:          config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountSocialUpdateHandler) Handle(ctx context.Context, command AccountSocialUpdateCommand) (*AccountSocialUpdateResult, *i18np.Error) {
	p, err := h.platformRepo.GetBySlug(ctx, command.Platform)
	if err != nil {
		return nil, err
	}
	err = h.platformFactory.ValidatePlatformValue(p, command.Value)
	if err != nil {
		return nil, err
	}
	social := &account.EntitySocial{
		Platform:   command.Platform,
		Value:      command.Value,
		FixedValue: h.platformFactory.FixPlatformValue(p, command.Value),
	}
	err = h.accountRepo.SocialUpdate(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, social)
	if err != nil {
		return nil, err
	}
	h.events.SocialUpdated(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, *social)
	return nil, nil
}
