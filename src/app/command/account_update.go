package command

import (
	"context"
	"time"

	"api.turistikrota.com/account/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountUpdateCommand struct {
	UserUUID    string
	UserName    string
	FullName    string
	Description string
	BirthDate   *time.Time
}

type AccountUpdateResult struct{}

type AccountUpdateHandler decorator.CommandHandler[AccountUpdateCommand, *AccountUpdateResult]

type accountUpdateHandler struct {
	repo    account.Repository
	factory account.Factory
	events  account.Events
}

type AccountUpdateHandlerConfig struct {
	Repo     account.Repository
	Factory  account.Factory
	Events   account.Events
	CqrsBase decorator.Base
}

func NewAccountUpdateHandler(config AccountUpdateHandlerConfig) AccountUpdateHandler {
	return decorator.ApplyCommandDecorators[AccountUpdateCommand, *AccountUpdateResult](
		accountUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountUpdateHandler) Handle(ctx context.Context, command AccountUpdateCommand) (*AccountUpdateResult, *i18np.Error) {
	u := account.UserUnique{
		UUID: command.UserUUID,
		Name: command.UserName,
	}
	acc, err := h.repo.Get(ctx, u)
	if err != nil {
		return nil, err
	}
	if acc.BirthDate != command.BirthDate {
		ageErr := h.factory.ValidateMinAge(command.BirthDate)
		if ageErr != nil {
			return nil, ageErr
		}
		acc.BirthDate = command.BirthDate
	}
	acc.FullName = command.FullName
	acc.Description = command.Description
	acc.CompletedRate = h.factory.CalcCompletedRate(acc)
	t := time.Now()
	acc.UpdatedAt = &t
	err = h.repo.Update(ctx, u, acc)
	if err != nil {
		return nil, err
	}
	h.events.Updated(u, *acc)
	return &AccountUpdateResult{}, nil
}
