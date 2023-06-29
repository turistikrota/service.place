package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountEnableCommand struct {
	UserUUID    string
	AccountName string
}

type AccountEnableResult struct{}

type AccountEnableHandler decorator.CommandHandler[AccountEnableCommand, *AccountEnableResult]

type accountEnableHandler struct {
	repo    account.Repository
	factory account.Factory
	events  account.Events
}

type AccountEnableHandlerConfig struct {
	Repo     account.Repository
	Factory  account.Factory
	Events   account.Events
	CqrsBase decorator.Base
}

func NewAccountEnableHandler(config AccountEnableHandlerConfig) AccountEnableHandler {
	return decorator.ApplyCommandDecorators[AccountEnableCommand, *AccountEnableResult](
		accountEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountEnableHandler) Handle(ctx context.Context, command AccountEnableCommand) (*AccountEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	if err != nil {
		return nil, err
	}
	h.events.Enabled(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	return nil, nil
}
