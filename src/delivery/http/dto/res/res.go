package res

import "github.com/turistikrota/service.place/src/app/query"

type Response interface {
	PlaceView(res *query.PlaceViewResult) *PlaceViewResponse
	FeatureList(res *query.FeatureListAllResult) []*FeatureListResponse
	PlaceList(res *query.PlaceFilterResult) *PlaceFilterResponse
}

type response struct{}

func New() Response {
	return &response{}
}
