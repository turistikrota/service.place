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
	PlaceFilterQuery struct {
		Filter place.EntityFilter
		Offset int64
		Limit  int64
	}
	PlaceFilterResult struct {
		Result list.Result[*place.Entity]
	}
	PlaceFilterHandler decorator.QueryHandler[PlaceFilterQuery, *PlaceFilterResult]
	placeFilterHandler struct {
		repo  place.Repository
		cache cache.Client[*list.Result[*place.Entity]]
	}
	PlaceFilterHandlerConfig struct {
		Repo     place.Repository
		CacheSrv cache.Service
		CqrsBase decorator.Base
	}
)

func NewPlaceFilterHandler(config PlaceFilterHandlerConfig) PlaceFilterHandler {
	return decorator.ApplyQueryDecorators[PlaceFilterQuery, *PlaceFilterResult](
		placeFilterHandler{
			repo:  config.Repo,
			cache: cache.New[*list.Result[*place.Entity]](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h placeFilterHandler) Handle(ctx context.Context, query PlaceFilterQuery) (*PlaceFilterResult, *i18np.Error) {
	if query.Filter.IsZero() {
		return h.withCache(ctx, query)
	}
	return h.withoutCache(ctx, query)
}

func (h placeFilterHandler) withCache(ctx context.Context, query PlaceFilterQuery) (*PlaceFilterResult, *i18np.Error) {
	cacheHandler := func() (*list.Result[*place.Entity], *i18np.Error) {
		return h.filter(ctx, query)
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &PlaceFilterResult{
		Result: *res,
	}, nil
}

func (h placeFilterHandler) withoutCache(ctx context.Context, query PlaceFilterQuery) (*PlaceFilterResult, *i18np.Error) {
	res, err := h.filter(ctx, query)
	if err != nil {
		return nil, err
	}
	return &PlaceFilterResult{
		Result: *res,
	}, nil
}

func (h placeFilterHandler) filter(ctx context.Context, query PlaceFilterQuery) (*list.Result[*place.Entity], *i18np.Error) {
	return h.repo.Filter(ctx, query.Filter, list.Config{
		Offset: query.Offset,
		Limit:  query.Limit,
	})
}

func (h placeFilterHandler) createCacheEntity() *list.Result[*place.Entity] {
	return &list.Result[*place.Entity]{
		List: make([]*place.Entity, 0),
	}
}

func (h placeFilterHandler) generateCacheKey(query PlaceFilterQuery) string {
	return fmt.Sprintf("place_filter_%v_%v", query.Offset, query.Limit)
}
