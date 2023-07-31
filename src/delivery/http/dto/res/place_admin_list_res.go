package res

import (
	"time"

	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/domain/place"
)

type PlaceAdminFilterResponse struct {
	List          []PlaceAdminFilterChild `json:"list"`
	Total         int64                   `json:"total"`
	FilteredTotal int64                   `json:"filteredTotal"`
	IsNext        bool                    `json:"isNext"`
	IsPrev        bool                    `json:"isPrev"`
	Page          int64                   `json:"page"`
}

type PlaceAdminFilterChild struct {
	UUID             string                        `json:"uuid"`
	FeatureUUIDs     []string                      `json:"featureUUIDs"`
	Images           []place.Image                 `json:"images"`
	Translations     map[place.Locale]Translations `json:"translations"`
	AverageTimeSpent place.TimeSpent               `json:"averageTimeSpent"`
	Review           place.Review                  `json:"review"`
	Coordinates      []float64                     `json:"coordinates"`
	IsActive         bool                          `json:"isActive"`
	IsDeleted        bool                          `json:"isDeleted"`
	IsPayed          bool                          `json:"isPayed"`
	Type             place.Type                    `json:"type"`
	UpdatedAt        time.Time                     `json:"updatedAt"`
	CreatedAt        time.Time                     `json:"createdAt"`
}

func (r *response) PlaceAdminList(res *query.PlaceAdminFilterResult) *PlaceAdminFilterResponse {
	return &PlaceAdminFilterResponse{
		List:          r.placeAdminListChild(res.Result.List),
		Total:         res.Result.Total,
		FilteredTotal: res.Result.FilteredTotal,
		IsNext:        res.Result.IsNext,
		IsPrev:        res.Result.IsPrev,
		Page:          res.Result.Page,
	}
}

func (r *response) placeAdminListChild(list []*place.Entity) []PlaceAdminFilterChild {
	res := make([]PlaceAdminFilterChild, len(list))
	for i, v := range list {
		res[i] = PlaceAdminFilterChild{
			UUID:             v.UUID,
			FeatureUUIDs:     v.FeatureUUIDs,
			Images:           v.Images,
			Translations:     r.translations(v.Translations),
			AverageTimeSpent: v.AverageTimeSpent,
			Review:           v.Review,
			Coordinates:      v.Coordinates,
			IsActive:         v.IsActive,
			IsDeleted:        v.IsDeleted,
			IsPayed:          v.IsPayed,
			Type:             v.Type,
			UpdatedAt:        v.UpdatedAt,
			CreatedAt:        v.CreatedAt,
		}
	}
	return res
}
