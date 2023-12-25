package query

import (
	"context"
	"fmt"

	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	PlaceAdminFilterQuery struct {
		AdminFilter place.EntityFilter
		Offset      int64
		Limit       int64
	}
	PlaceAdminFilterResult struct {
		Result list.Result[*place.Entity]
	}
	PlaceAdminFilterHandler decorator.QueryHandler[PlaceAdminFilterQuery, *PlaceAdminFilterResult]
	placeAdminFilterHandler struct {
		repo  place.Repository
		cache cache.Client[*list.Result[*place.Entity]]
	}
	PlaceAdminFilterHandlerConfig struct {
		Repo     place.Repository
		CacheSrv cache.Service
		CqrsBase decorator.Base
	}
)

func NewPlaceAdminFilterHandler(config PlaceAdminFilterHandlerConfig) PlaceAdminFilterHandler {
	return decorator.ApplyQueryDecorators[PlaceAdminFilterQuery, *PlaceAdminFilterResult](
		placeAdminFilterHandler{
			repo:  config.Repo,
			cache: cache.New[*list.Result[*place.Entity]](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h placeAdminFilterHandler) Handle(ctx context.Context, query PlaceAdminFilterQuery) (*PlaceAdminFilterResult, *i18np.Error) {
	if query.AdminFilter.IsZero() {
		return h.withCache(ctx, query)
	}
	return h.withoutCache(ctx, query)
}

func (h placeAdminFilterHandler) withCache(ctx context.Context, query PlaceAdminFilterQuery) (*PlaceAdminFilterResult, *i18np.Error) {
	cacheHandler := func() (*list.Result[*place.Entity], *i18np.Error) {
		return h.filter(ctx, query)
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &PlaceAdminFilterResult{
		Result: *res,
	}, nil
}

func (h placeAdminFilterHandler) withoutCache(ctx context.Context, query PlaceAdminFilterQuery) (*PlaceAdminFilterResult, *i18np.Error) {
	res, err := h.filter(ctx, query)
	if err != nil {
		return nil, err
	}
	return &PlaceAdminFilterResult{
		Result: *res,
	}, nil
}

func (h placeAdminFilterHandler) filter(ctx context.Context, query PlaceAdminFilterQuery) (*list.Result[*place.Entity], *i18np.Error) {
	return h.repo.AdminList(ctx, query.AdminFilter, list.Config{
		Offset: query.Offset,
		Limit:  query.Limit,
	})
}

func (h placeAdminFilterHandler) createCacheEntity() *list.Result[*place.Entity] {
	return &list.Result[*place.Entity]{
		List: make([]*place.Entity, 0),
	}
}

func (h placeAdminFilterHandler) generateCacheKey(query PlaceAdminFilterQuery) string {
	return fmt.Sprintf("place_admin_list_%v_%v", query.Offset, query.Limit)
}
