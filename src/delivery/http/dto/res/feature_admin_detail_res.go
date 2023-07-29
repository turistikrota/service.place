package res

import (
	"time"

	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type AdminFeatureDetailResponse struct {
	UUID         string                                  `json:"uuid"`
	Icon         string                                  `json:"icon"`
	Translations map[feature.Locale]feature.Translations `json:"translations"`
	IsActive     bool                                    `json:"isActive"`
	IsDeleted    bool                                    `json:"isDeleted"`
	UpdatedAt    time.Time                               `json:"updatedAt"`
	CreatedAt    time.Time                               `json:"createdAt"`
}

func (r *response) AdminFeatureDetail(res *query.AdminFeatureDetailResult) *AdminFeatureDetailResponse {
	return &AdminFeatureDetailResponse{
		UUID:         res.Feature.UUID,
		Icon:         res.Feature.Icon,
		Translations: res.Feature.Translations,
		IsActive:     res.Feature.IsActive,
		IsDeleted:    res.Feature.IsDeleted,
		UpdatedAt:    res.Feature.UpdatedAt,
		CreatedAt:    res.Feature.CreatedAt,
	}
}
