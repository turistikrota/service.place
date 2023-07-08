package res

import (
	"time"

	"github.com/turistikrota/service.place/src/app/query"
	"github.com/turistikrota/service.place/src/domain/place"
)

type PlaceViewResponse struct {
	Features         []query.PlaceViewFeatureItem        `json:"features"`
	Images           []place.Image                       `json:"images"`
	Translations     map[place.Locale]place.Translations `json:"translations"`
	AverageTimeSpent place.TimeSpent                     `json:"averageTimeSpent"`
	Review           place.Review                        `json:"review"`
	Coordinates      []float64                           `json:"coordinates"`
	IsPayed          bool                                `json:"isPayed"`
	Type             place.Type                          `json:"type"`
	CreatedAt        time.Time                           `json:"createdAt"`
}

func (r *response) PlaceView(res *query.PlaceViewResult) *PlaceViewResponse {
	return &PlaceViewResponse{
		Features:         res.Features,
		Images:           res.Place.Images,
		Translations:     res.Place.Translations,
		AverageTimeSpent: res.Place.AverageTimeSpent,
		Review:           res.Place.Review,
		Coordinates:      res.Place.Coordinates,
		IsPayed:          res.Place.IsPayed,
		Type:             res.Place.Type,
		CreatedAt:        res.Place.CreatedAt,
	}
}
