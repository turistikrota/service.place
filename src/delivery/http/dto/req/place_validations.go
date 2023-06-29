package req

import "api.turistikrota.com/place/src/domain/place"

type PlaceImageRequest struct {
	Url   string `json:"url" validate:"required,url"`
	Order *int16 `json:"order" validate:"required,min=0,max=20"`
}

type PlaceTimeSpentRequest struct {
	Min int16 `json:"min" validate:"required,min=0,max=200"`
	Max int16 `json:"max" validate:"required,min=0,max=200"`
}

type PlaceTranslationRequest struct {
	Title       string                     `json:"title" validate:"required,min=1,max=255"`
	Description string                     `json:"description" validate:"required,min=1,max=255"`
	Locale      string                     `json:"locale" validate:"required,min=1,max=255"`
	Slug        string                     `json:"slug" validate:"required,min=1,max=255"`
	MarkdownURL string                     `json:"markdown_url" validate:"required,min=1,max=255"`
	Seo         PlaceTranslationSeoRequest `json:"seo" validate:"required,dive,required"`
}

type PlaceTranslationSeoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1,max=255"`
	Keywords    string `json:"keywords" validate:"required,min=1,max=255"`
}

func (r *PlaceTranslationSeoRequest) ToSeo() place.Seo {
	return place.Seo{
		Title:       r.Title,
		Description: r.Description,
		Keywords:    r.Keywords,
	}
}

func (r *PlaceTimeSpentRequest) ToTimeSpent() place.TimeSpent {
	return place.TimeSpent{
		Min: r.Min,
		Max: r.Max,
	}
}
