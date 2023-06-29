package query

import (
	"context"

	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformGetBySlugQuery struct {
	Slug string
}

type PlatformGetBySlugResult struct {
	Entity platform.Entity
}

type PlatformGetBySlugHandler decorator.QueryHandler[PlatformGetBySlugQuery, *PlatformGetBySlugResult]

type platformGetBySlugHandler struct {
	repo platform.Repository
}

type PlatformGetBySlugHandlerConfig struct {
	Repo     platform.Repository
	CqrsBase decorator.Base
}

func NewPlatformGetBySlugHandler(config PlatformGetBySlugHandlerConfig) PlatformGetBySlugHandler {
	return decorator.ApplyQueryDecorators[PlatformGetBySlugQuery, *PlatformGetBySlugResult](
		platformGetBySlugHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h platformGetBySlugHandler) Handle(ctx context.Context, query PlatformGetBySlugQuery) (*PlatformGetBySlugResult, *i18np.Error) {
	a, err := h.repo.GetBySlug(ctx, query.Slug)
	if err != nil {
		return nil, err
	}
	return &PlatformGetBySlugResult{
		Entity: *a,
	}, nil
}
