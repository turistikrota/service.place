package query

import (
	"context"
	"fmt"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceFilterQuery   struct{}
	PlaceFilterResult  struct{}
	PlaceFilterHandler decorator.QueryHandler[PlaceFilterQuery, *PlaceFilterResult]
	placeFilterHandler struct {
		repo  place.Repository
		cache cache.Client[*place.Entity]
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
			cache: cache.New[*place.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h placeFilterHandler) Handle(ctx context.Context, query PlaceFilterQuery) (*PlaceFilterResult, *i18np.Error) {
	/*
		    cacheHandler := func() (*post.Entity, *i18np.Error) {
				return h.repo.Filter(ctx, post.I18nDetail{
					Locale: query.Locale,
					Slug:   query.Slug,
				})
			}
			res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(h.generateCacheKey(query))
		    if err != nil {
				return nil, err
			}
	*/
	return &PlaceFilterResult{}, nil
}

func (h placeFilterHandler) createCacheEntity() *place.Entity {
	return &place.Entity{}
}

func (h placeFilterHandler) generateCacheKey(query PlaceFilterQuery) string {
	return fmt.Sprintf("cache_key")
}
