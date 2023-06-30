package query

import (
	"context"
	"fmt"

	"api.turistikrota.com/place/src/domain/feature"
	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceViewQuery struct {
		Locale string
		Slug   string
	}
	PlaceViewResult struct {
		Place    *place.Entity
		Features []PlaceViewFeatureItem
	}
	PlaceViewHandler decorator.QueryHandler[PlaceViewQuery, *PlaceViewResult]
	placeViewHandler struct {
		repo        place.Repository
		featureRepo feature.Repository
		cache       cache.Client[*place.Entity]
	}
	PlaceViewHandlerConfig struct {
		Repo        place.Repository
		FeatureRepo feature.Repository
		CacheSrv    cache.Service
		CqrsBase    decorator.Base
	}
	PlaceViewFeatureItem struct {
		UUID         string                                  `json:"uuid"`
		Icon         string                                  `json:"icon"`
		Translations map[feature.Locale]feature.Translations `json:"translations"`
	}
)

func NewPlaceViewHandler(config PlaceViewHandlerConfig) PlaceViewHandler {
	return decorator.ApplyQueryDecorators[PlaceViewQuery, *PlaceViewResult](
		placeViewHandler{
			repo:        config.Repo,
			featureRepo: config.FeatureRepo,
			cache:       cache.New[*place.Entity](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h placeViewHandler) Handle(ctx context.Context, query PlaceViewQuery) (*PlaceViewResult, *i18np.Error) {
	cacheHandler := func() (*place.Entity, *i18np.Error) {
		return h.repo.View(ctx, place.I18nDetail{
			Locale: query.Locale,
			Slug:   query.Slug,
		})
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	features, err := h.getFeatures(ctx, res.FeatureUUIDs)
	if err != nil {
		return nil, err
	}
	return &PlaceViewResult{
		Place:    res,
		Features: features,
	}, nil
}

func (h placeViewHandler) createCacheEntity() *place.Entity {
	return &place.Entity{}
}

func (h placeViewHandler) generateCacheKey(query PlaceViewQuery) string {
	return fmt.Sprintf("place_view_%s_%s", query.Locale, query.Slug)
}

func (h placeViewHandler) getFeatures(ctx context.Context, uuids []string) ([]PlaceViewFeatureItem, *i18np.Error) {
	res, err := h.featureRepo.GetByUUIDs(ctx, uuids)
	if err != nil {
		return nil, err
	}
	list := make([]PlaceViewFeatureItem, len(res))
	for i, item := range res {
		list[i] = PlaceViewFeatureItem{
			UUID:         item.UUID,
			Icon:         item.Icon,
			Translations: item.Translations,
		}
	}
	return list, nil
}
