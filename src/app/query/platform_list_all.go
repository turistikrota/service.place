package query

import (
	"context"

	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformListAllQuery struct{}

type PlatformListAllResult struct {
	Entities []*platform.Entity
}

type PlatformListAllHandler decorator.QueryHandler[PlatformListAllQuery, *PlatformListAllResult]

type platformListAllHandler struct {
	repo platform.Repository
}

type PlatformListAllHandlerConfig struct {
	Repo     platform.Repository
	CqrsBase decorator.Base
}

func NewPlatformListAllHandler(config PlatformListAllHandlerConfig) PlatformListAllHandler {
	return decorator.ApplyQueryDecorators[PlatformListAllQuery, *PlatformListAllResult](
		platformListAllHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h platformListAllHandler) Handle(ctx context.Context, query PlatformListAllQuery) (*PlatformListAllResult, *i18np.Error) {
	a, err := h.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return &PlatformListAllResult{
		Entities: a,
	}, nil
}
