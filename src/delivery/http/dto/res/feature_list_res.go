package res

import (
	"api.turistikrota.com/place/src/app/query"
	"api.turistikrota.com/place/src/domain/feature"
)

type FeatureListResponse struct {
	UUID         string                                  `json:"uuid"`
	Icon         string                                  `json:"icon"`
	Translations map[feature.Locale]feature.Translations `json:"translations"`
}

func (r *response) FeatureList(res *query.FeatureListAllResult) []*FeatureListResponse {
	list := make([]*FeatureListResponse, len(res.Features))
	for i, v := range res.Features {
		list[i] = &FeatureListResponse{
			UUID:         v.UUID,
			Icon:         v.Icon,
			Translations: v.Translations,
		}
	}
	return list
}
