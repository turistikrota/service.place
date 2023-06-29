package req

import "api.turistikrota.com/place/src/app/query"

type PlatformGetRequest struct {
	Slug string `param:"slug" validate:"required"`
}

func (r *PlatformGetRequest) ToQuery() query.PlatformGetBySlugQuery {
	return query.PlatformGetBySlugQuery{
		Slug: r.Slug,
	}
}
