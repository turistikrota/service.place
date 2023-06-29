package query

import (
	"context"

	"api.turistikrota.com/account/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountProfileViewQuery struct {
	Name string
}

type AccountProfileViewResult struct {
	Entity account.Entity
}

type AccountProfileViewHandler decorator.QueryHandler[AccountProfileViewQuery, *AccountProfileViewResult]

type accountProfileViewHandler struct {
	repo account.Repository
}

type AccountProfileViewHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountProfileViewHandler(config AccountProfileViewHandlerConfig) AccountProfileViewHandler {
	return decorator.ApplyQueryDecorators[AccountProfileViewQuery, *AccountProfileViewResult](
		accountProfileViewHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h accountProfileViewHandler) Handle(ctx context.Context, query AccountProfileViewQuery) (*AccountProfileViewResult, *i18np.Error) {
	a, err := h.repo.ProfileView(ctx, account.UserUnique{
		Name: query.Name,
	})
	if err != nil {
		return nil, err
	}
	return &AccountProfileViewResult{
		Entity: *a,
	}, nil
}
