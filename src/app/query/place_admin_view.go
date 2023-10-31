package query

import (
	"context"

	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	AdminPlaceViewQuery struct {
		UUID string
	}
	AdminPlaceViewResult struct {
		Place    *place.Entity
		Features []AdminPlaceViewFeatureItem
	}
	AdminPlaceViewHandler decorator.QueryHandler[AdminPlaceViewQuery, *AdminPlaceViewResult]
	adminPlaceViewHandler struct {
		repo        place.Repository
		featureRepo feature.Repository
		cdnUrl      string
	}
	AdminPlaceViewHandlerConfig struct {
		Repo        place.Repository
		FeatureRepo feature.Repository
		CacheSrv    cache.Service
		CqrsBase    decorator.Base
		CdnUrl      string
	}
	AdminPlaceViewFeatureItem struct {
		UUID         string                                  `json:"uuid"`
		Icon         string                                  `json:"icon"`
		Translations map[feature.Locale]feature.Translations `json:"translations"`
	}
)

func NewAdminPlaceViewHandler(config AdminPlaceViewHandlerConfig) AdminPlaceViewHandler {
	return decorator.ApplyQueryDecorators[AdminPlaceViewQuery, *AdminPlaceViewResult](
		adminPlaceViewHandler{
			repo:        config.Repo,
			featureRepo: config.FeatureRepo,
			cdnUrl:      config.CdnUrl,
		},
		config.CqrsBase,
	)
}

func (h adminPlaceViewHandler) Handle(ctx context.Context, query AdminPlaceViewQuery) (*AdminPlaceViewResult, *i18np.Error) {
	res, err := h.repo.AdminView(ctx, query.UUID)
	if err != nil {
		return nil, err
	}
	trTranslations, trOk := res.Translations[place.LocaleTR]
	enTranslations, enOk := res.Translations[place.LocaleEN]
	if enOk && enTranslations.MarkdownURL == "" {
		enTranslations.MarkdownURL = dressCdnMarkdown(h.cdnUrl, res.UUID, place.LocaleEN.String())
	}
	if trOk && trTranslations.MarkdownURL == "" {
		trTranslations.MarkdownURL = dressCdnMarkdown(h.cdnUrl, res.UUID, place.LocaleTR.String())
		
	}
	features, err := h.getFeatures(ctx, res.FeatureUUIDs)
	if err != nil {
		return nil, err
	}
	return &AdminPlaceViewResult{
		Place:    res,
		Features: features,
	}, nil
}

func (h adminPlaceViewHandler) getFeatures(ctx context.Context, uuids []string) ([]AdminPlaceViewFeatureItem, *i18np.Error) {
	res, err := h.featureRepo.GetByUUIDs(ctx, uuids)
	if err != nil {
		return nil, err
	}
	list := make([]AdminPlaceViewFeatureItem, len(res))
	for i, item := range res {
		list[i] = AdminPlaceViewFeatureItem{
			UUID:         item.UUID,
			Icon:         item.Icon,
			Translations: item.Translations,
		}
	}
	return list, nil
}
