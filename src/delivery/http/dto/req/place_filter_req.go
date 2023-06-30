package req

import (
	"api.turistikrota.com/place/src/app/query"
	"api.turistikrota.com/place/src/domain/place"
)

type PlaceFilterRequest struct {
	Page            int64
	Limit           int64
	Query           string                       `json:"query,omitempty" validate:"omitempty,max=100"`
	Coordinates     []float64                    `json:"coordinates,omitempty" validate:"omitempty,min=2,max=2"`
	FeatureUUIDs    []string                     `json:"feature_uuids,omitempty" validate:"omitempty,min=1,max=10,dive,object_id"`
	TimeSpent       *PlaceFilterTimeSpentRequest `json:"time_spent,omitempty" validate:"omitempty,dive,required"`
	Distance        float64                      `json:"distance,omitempty" validate:"omitempty,gt=0"`
	IsPayed         *bool                        `json:"is_payed,omitempty" validate:"omitempty"`
	MinReview       *int16                       `json:"min_review,omitempty" validate:"omitempty,gt=0"`
	MaxReview       *int16                       `json:"max_review,omitempty" validate:"omitempty,gt=0"`
	MinAveragePoint *float32                     `json:"min_average_point,omitempty" validate:"omitempty,gt=0"`
	MaxAveragePoint *float32                     `json:"max_average_point,omitempty" validate:"omitempty,gt=0"`
}

type PlaceFilterTimeSpentRequest struct {
	Min int16 `json:"min" validate:"required,gt=0"`
	Max int16 `json:"max" validate:"required,gt=0"`
}

func (r *PlaceFilterRequest) LoadPagination(pagination *PaginationRequest) *PlaceFilterRequest {
	pagination.Default()
	r.Page = *pagination.Page
	r.Limit = *pagination.Limit
	return r
}

func (r *PlaceFilterRequest) ToQuery(locale string) query.PlaceFilterQuery {
	return query.PlaceFilterQuery{
		Filter: place.EntityFilter{
			Locale:           locale,
			Query:            r.Query,
			Coordinates:      r.Coordinates,
			FeatureUUIDs:     r.FeatureUUIDs,
			AverageTimeSpent: r.toAverageTimeSpent(),
			Distance:         r.Distance,
			IsPayed:          r.IsPayed,
			MinAveragePoint:  r.MinAveragePoint,
			MinReview:        r.MinReview,
			MaxReview:        r.MaxReview,
			MaxAveragePoint:  r.MaxAveragePoint,
		},
		Offset: r.Page * r.Limit,
		Limit:  r.Limit,
	}
}

func (r *PlaceFilterRequest) toAverageTimeSpent() *place.TimeSpent {
	if r.TimeSpent == nil {
		return nil
	}
	return &place.TimeSpent{
		Min: r.TimeSpent.Min,
		Max: r.TimeSpent.Max,
	}
}