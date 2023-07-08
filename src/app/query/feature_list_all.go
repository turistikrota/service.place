package query

import (
	"context"

	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type (
	FeatureListAllQuery  struct{}
	FeatureListAllResult struct {
		Features []*feature.Entity
	}
	FeatureListAllHandler decorator.QueryHandler[FeatureListAllQuery, *FeatureListAllResult]
	featureListAllHandler struct {
		repo  feature.Repository
		cache cache.Client[[]*feature.Entity]
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
			cache: cache.New[[]*feature.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h featureListAllHandler) Handle(ctx context.Context, query FeatureListAllQuery) (*FeatureListAllResult, *i18np.Error) {
	cacheHandler := func() ([]*feature.Entity, *i18np.Error) {
		return h.repo.ListAll(ctx)
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &FeatureListAllResult{
		Features: res,
	}, nil
}

func (h featureListAllHandler) createCacheEntity() []*feature.Entity {
	return make([]*feature.Entity, 0)
}

func (h featureListAllHandler) generateCacheKey(query FeatureListAllQuery) string {
	return "feature_list_all"
}
