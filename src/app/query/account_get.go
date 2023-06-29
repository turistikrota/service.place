package query

import (
	"context"
	"fmt"
	"time"

	"api.turistikrota.com/place/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/cache"
	"github.com/turistikrota/service.shared/db/redis"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountGetQuery struct {
	UserUUID string
	Name     string
}

type AccountGetResult struct {
	Entity account.Entity
}

type AccountGetHandler decorator.QueryHandler[AccountGetQuery, *AccountGetResult]

type accountGetHandler struct {
	repo  account.Repository
	cache cache.Client[*account.Entity]
}

type AccountGetHandlerConfig struct {
	Repo     account.Repository
	CacheSrv redis.Service
	CqrsBase decorator.Base
}

func NewAccountGetHandler(config AccountGetHandlerConfig) AccountGetHandler {
	return decorator.ApplyQueryDecorators[AccountGetQuery, *AccountGetResult](
		accountGetHandler{
			repo:  config.Repo,
			cache: cache.New[*account.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h accountGetHandler) Handle(ctx context.Context, query AccountGetQuery) (*AccountGetResult, *i18np.Error) {
	creator := func() *account.Entity {
		return &account.Entity{}
	}
	cacheHandler := func() (*account.Entity, *i18np.Error) {
		return h.repo.Get(ctx, account.UserUnique{
			UUID: query.UserUUID,
			Name: query.Name,
		})
	}
	a, err := h.cache.Creator(creator).Handler(cacheHandler).Timeout(1 * time.Minute).Get(h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &AccountGetResult{
		Entity: *a,
	}, nil
}

func (h accountGetHandler) generateCacheKey(query AccountGetQuery) string {
	return fmt.Sprintf("c_acc_get__%v_%v", query.Name, query.UserUUID)
}
