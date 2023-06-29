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
	FeatureListAllQuery   struct{}
	FeatureListAllResult  struct{}
	FeatureListAllHandler decorator.QueryHandler[FeatureListAllQuery, *FeatureListAllResult]
	featureListAllHandler struct {
		repo  feature.Repository
		cache cache.Client[*feature.Entity]
	}
	FeatureListAllHandlerConfig struct {
		Repo     feature.Repository
		CacheSrv cache.Service
		CqrsBase decorator.Base
	}
)

func NewFeatureListAllHandler(config FeatureListAllHandlerConfig) FeatureListAllHandler {
	return decorator.ApplyQueryDecorators[FeatureListAllQuery, *FeatureListAllResult](
		featureListAllHandler{
			repo:  config.Repo,
			cache: cache.New[*feature.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h featureListAllHandler) Handle(ctx context.Context, query FeatureListAllQuery) (*FeatureListAllResult, *i18np.Error) {
	/*
		    cacheHandler := func() (*post.Entity, *i18np.Error) {
				return h.repo.ListAll(ctx, post.I18nDetail{
					Locale: query.Locale,
					Slug:   query.Slug,
				})
			}
			res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(h.generateCacheKey(query))
		    if err != nil {
				return nil, err
			}
	*/
	return &FeatureListAllResult{}, nil
}

func (h featureListAllHandler) createCacheEntity() *feature.Entity {
	return &feature.Entity{}
}

func (h featureListAllHandler) generateCacheKey(query FeatureListAllQuery) string {
	return fmt.Sprintf("cache_key")
}
