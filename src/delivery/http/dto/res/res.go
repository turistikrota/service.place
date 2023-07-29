package res

import "github.com/turistikrota/service.place/src/app/query"

type Response interface {
	PlaceView(res *query.PlaceViewResult) *PlaceViewResponse
	FeatureList(res *query.FeatureListAllResult) []*FeatureListResponse
	PlaceList(res *query.PlaceFilterResult) *PlaceFilterResponse
	AdminFeatureList(res *query.AdminFeatureListAllResult) []*AdminFeatureListResponse
	AdminFeatureDetail(res *query.AdminFeatureDetailResult) *AdminFeatureDetailResponse
}

type response struct{}

func New() Response {
	return &response{}
}
