package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountCreateCommand struct {
	UserUUID    string
	AccountName string
}

type AccountCreateResult struct {
	AccountName string
}

type AccountCreateHandler decorator.CommandHandler[AccountCreateCommand, *AccountCreateResult]

type accountCreateHandler struct {
	repo    account.Repository
	factory account.Factory
	events  account.Events
}

type AccountCreateHandlerConfig struct {
	Repo     account.Repository
	Factory  account.Factory
	Events   account.Events
	CqrsBase decorator.Base
}

func NewAccountCreateHandler(config AccountCreateHandlerConfig) AccountCreateHandler {
	return decorator.ApplyCommandDecorators[AccountCreateCommand, *AccountCreateResult](
		accountCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountCreateHandler) Handle(ctx context.Context, command AccountCreateCommand) (*AccountCreateResult, *i18np.Error) {
	u := account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}
	exist, err := h.repo.Exist(ctx, u)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, h.factory.Errors.ErrAlreadyExist()
	}
	acc := h.factory.NewAccount(command.UserUUID, command.AccountName)
	err = h.factory.Validate(acc)
	if err != nil {
		return nil, err
	}
	acc, err = h.repo.Create(ctx, acc)
	if err != nil {
		return nil, err
	}
	h.events.Created(u)
	return &AccountCreateResult{
		AccountName: acc.UserName,
	}, nil
}
