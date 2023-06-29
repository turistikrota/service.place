package res

import (
	"time"

	"api.turistikrota.com/place/src/app/query"
	"api.turistikrota.com/place/src/domain/platform"
)

type PlatformGetResponse struct {
	Name         string                                          `json:"name"`
	Slug         string                                          `json:"slug"`
	Regexp       string                                          `json:"regexp"`
	Prefix       string                                          `json:"prefix"`
	IsActive     bool                                            `json:"isActive"`
	Translations map[platform.Locale]PlatformTranslationResponse `json:"translations"`
	CreatedAt    time.Time                                       `json:"createdAt"`
}

type PlatformTranslationResponse struct {
	Name        string `json:"name"`
	Placeholder string `json:"placeholder"`
	Description string `json:"description"`
}

func (r *response) PlatformGet(res *query.PlatformGetBySlugResult) *PlatformGetResponse {
	return &PlatformGetResponse{
		Name:         res.Entity.Name,
		Slug:         res.Entity.Slug,
		Regexp:       res.Entity.Regexp,
		Prefix:       res.Entity.Prefix,
		Translations: r.PlatformTranslations(res.Entity.Translations),
		IsActive:     res.Entity.IsActive,
		CreatedAt:    res.Entity.CreatedAt,
	}
}

func (r *response) PlatformTranslations(translations map[platform.Locale]platform.Translations) map[platform.Locale]PlatformTranslationResponse {
	res := make(map[platform.Locale]PlatformTranslationResponse, 0)
	for i, t := range translations {
		res[i] = PlatformTranslationResponse{
			Name:        t.Name,
			Placeholder: t.Placeholder,
			Description: t.Description,
		}
	}
	return res
}
