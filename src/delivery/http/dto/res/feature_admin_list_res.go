package res

import (
	"time"

	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type AdminFeatureListResponse struct {
	UUID         string                                  `json:"uuid"`
	Icon         string                                  `json:"icon"`
	Translations map[feature.Locale]feature.Translations `json:"translations"`
	IsActive     bool                                    `json:"isActive"`
	IsDeleted    bool                                    `json:"isDeleted"`
	UpdatedAt    time.Time                               `json:"updatedAt"`
	CreatedAt    time.Time                               `json:"createdAt"`
}

func (r *response) AdminFeatureList(res *query.AdminFeatureListAllResult) []*AdminFeatureListResponse {
	list := make([]*AdminFeatureListResponse, len(res.Features))
	for i, v := range res.Features {
		list[i] = &AdminFeatureListResponse{
			UUID:         v.UUID,
			Icon:         v.Icon,
			Translations: v.Translations,
			IsActive:     v.IsActive,
			IsDeleted:    v.IsDeleted,
			UpdatedAt:    v.UpdatedAt,
			CreatedAt:    v.CreatedAt,
		}
	}
	return list
}
