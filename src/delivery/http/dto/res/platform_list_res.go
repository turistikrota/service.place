package res

import (
	"api.turistikrota.com/place/src/app/query"
	"api.turistikrota.com/place/src/domain/platform"
)

type PlatformListResponse struct {
	Name         string                                          `json:"name"`
	Slug         string                                          `json:"slug"`
	Regexp       string                                          `json:"regexp"`
	Prefix       string                                          `json:"prefix"`
	Translations map[platform.Locale]PlatformTranslationResponse `json:"translations"`
}

func (r *response) PlatformList(res *query.PlatformListAllResult) []PlatformListResponse {
	resList := make([]PlatformListResponse, 0)
	for _, v := range res.Entities {
		resList = append(resList, PlatformListResponse{
			Name:         v.Name,
			Slug:         v.Slug,
			Regexp:       v.Regexp,
			Prefix:       v.Prefix,
			Translations: r.PlatformTranslations(v.Translations),
		})
	}
	return resList
}
