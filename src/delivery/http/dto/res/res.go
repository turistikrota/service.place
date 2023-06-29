package res

import (
	"api.turistikrota.com/place/src/app/command"
	"api.turistikrota.com/place/src/app/query"
)

type Response interface {
	AccountCreate(res *command.AccountCreateResult) *AccountCreateResponse
	AccountGet(res *query.AccountGetResult) *AccountGetResponse
	AccountListMy(res *query.AccountListMyResult) []*AccountListMyResponse
	AccountProfileView(res *query.AccountProfileViewResult) *AccountProfileViewResponse
	PlatformGet(res *query.PlatformGetBySlugResult) *PlatformGetResponse
	PlatformList(res *query.PlatformListAllResult) []PlatformListResponse
}

type response struct{}

func New() Response {
	return &response{}
}
