package query

import (
	"context"
	"fmt"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	FeatureGetByUUIDsQuery   struct{}
	FeatureGetByUUIDsResult  struct{}
	FeatureGetByUUIDsHandler decorator.QueryHandler[FeatureGetByUUIDsQuery, *FeatureGetByUUIDsResult]
	featureGetByUUIDsHandler struct {
		repo  feature.Repository
		cache cache.Client[*feature.Entity]
	}
	FeatureGetByUUIDsHandlerConfig struct {
		Repo     feature.Repository
		CacheSrv cache.Service
		CqrsBase decorator.Base
	}
)

func NewFeatureGetByUUIDsHandler(config FeatureGetByUUIDsHandlerConfig) FeatureGetByUUIDsHandler {
	return decorator.ApplyQueryDecorators[FeatureGetByUUIDsQuery, *FeatureGetByUUIDsResult](
		featureGetByUUIDsHandler{
			repo:  config.Repo,
			cache: cache.New[*feature.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h featureGetByUUIDsHandler) Handle(ctx context.Context, query FeatureGetByUUIDsQuery) (*FeatureGetByUUIDsResult, *i18np.Error) {
	/*
		    cacheHandler := func() (*post.Entity, *i18np.Error) {
				return h.repo.GetByUUIDs(ctx, post.I18nDetail{
					Locale: query.Locale,
					Slug:   query.Slug,
				})
			}
			res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(h.generateCacheKey(query))
		    if err != nil {
				return nil, err
			}
	*/
	return &FeatureGetByUUIDsResult{}, nil
}

func (h featureGetByUUIDsHandler) createCacheEntity() *feature.Entity {
	return &feature.Entity{}
}

func (h featureGetByUUIDsHandler) generateCacheKey(query FeatureGetByUUIDsQuery) string {
	return fmt.Sprintf("cache_key")
}
