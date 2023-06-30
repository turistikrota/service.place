package req

import "api.turistikrota.com/place/src/app/query"

type PlaceViewRequest struct {
	Slug string `param:"slug" validate:"required,slug"`
}

func (r *PlaceViewRequest) ToQuery(locale string) query.PlaceViewQuery {
	return query.PlaceViewQuery{
		Slug:   r.Slug,
		Locale: locale,
	}
}
