package res

import (
	"time"

	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/domain/place"
)

type PlaceAdminViewResponse struct {
	UUID             string                               `json:"uuid"`
	Features         []query.AdminPlaceViewFeatureItem    `json:"features"`
	Images           []place.Image                        `json:"images"`
	Translations     map[place.Locale]*place.Translations `json:"translations"`
	Restorations     []place.Restoration                  `json:"restorations"`
	AverageTimeSpent place.TimeSpent                      `json:"averageTimeSpent"`
	Review           place.Review                         `json:"review"`
	Coordinates      []float64                            `json:"coordinates"`
	IsActive         bool                                 `json:"isActive"`
	IsDeleted        bool                                 `json:"isDeleted"`
	IsPayed          bool                                 `json:"isPayed"`
	Type             place.Type                           `json:"type"`
	UpdatedAt        time.Time                            `json:"updatedAt"`
	CreatedAt        time.Time                            `json:"createdAt"`
}

func (r *response) PlaceAdminView(res *query.AdminPlaceViewResult) *PlaceAdminViewResponse {
	return &PlaceAdminViewResponse{
		UUID:             res.Place.UUID,
		Features:         res.Features,
		Images:           res.Place.Images,
		Translations:     res.Place.Translations,
		Restorations:     res.Place.Restorations,
		AverageTimeSpent: res.Place.AverageTimeSpent,
		Review:           res.Place.Review,
		Coordinates:      res.Place.Coordinates,
		IsPayed:          res.Place.IsPayed,
		IsActive:         res.Place.IsActive,
		IsDeleted:        res.Place.IsDeleted,
		Type:             res.Place.Type,
		UpdatedAt:        res.Place.UpdatedAt,
		CreatedAt:        res.Place.CreatedAt,
	}
}
