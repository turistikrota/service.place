package res

import (
	"api.turistikrota.com/place/src/app/query"
	"api.turistikrota.com/place/src/domain/place"
)

type PlaceFilterResponse struct {
	List          []PlaceFilterChild `json:"list"`
	Total         int64              `json:"total"`
	FilteredTotal int64              `json:"filteredTotal"`
	IsNext        bool               `json:"isNext"`
	IsPrev        bool               `json:"isPrev"`
	Page          int64              `json:"page"`
}

type PlaceFilterChild struct {
	Images           []place.Image                 `json:"images"`
	Translations     map[place.Locale]Translations `json:"translations"`
	AverageTimeSpent place.TimeSpent               `json:"averageTimeSpent"`
	Review           place.Review                  `json:"review"`
	Coordinates      []float64                     `json:"coordinates"`
	IsPayed          bool                          `json:"isPayed"`
}

type Translations struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

func (r *response) PlaceList(res *query.PlaceFilterResult) *PlaceFilterResponse {
	return &PlaceFilterResponse{
		List:          r.placeListChild(res.Result.List),
		Total:         res.Result.Total,
		FilteredTotal: res.Result.FilteredTotal,
		IsNext:        res.Result.IsNext,
		IsPrev:        res.Result.IsPrev,
		Page:          res.Result.Page,
	}
}

func (r *response) placeListChild(list []*place.Entity) []PlaceFilterChild {
	res := make([]PlaceFilterChild, len(list))
	for i, v := range list {
		res[i] = PlaceFilterChild{
			Images:           v.Images,
			Translations:     r.translations(v.Translations),
			AverageTimeSpent: v.AverageTimeSpent,
			Review:           v.Review,
			Coordinates:      v.Coordinates,
			IsPayed:          v.IsPayed,
		}
	}
	return res
}

func (r *response) translations(translations map[place.Locale]place.Translations) map[place.Locale]Translations {
	res := make(map[place.Locale]Translations, len(translations))
	for k, v := range translations {
		res[k] = Translations{
			Title:       v.Title,
			Description: v.Description,
			Slug:        v.Slug,
		}
	}
	return res
}
