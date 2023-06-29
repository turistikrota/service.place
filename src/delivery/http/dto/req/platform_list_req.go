package req

import "api.turistikrota.com/account/src/app/query"

type PlatformListRequest struct{}

func (r *PlatformListRequest) ToQuery() query.PlatformListAllQuery {
	return query.PlatformListAllQuery{}
}
