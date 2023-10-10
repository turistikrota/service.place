package req

import "github.com/turistikrota/service.place/src/domain/place"

type PlaceImageRequest struct {
	Url   string `json:"url" validate:"required,url"`
	Order *int16 `json:"order" validate:"required,min=0,max=200"`
}

type PlaceTimeSpentRequest struct {
	Min int16 `json:"min" validate:"required,min=0,max=2000"`
	Max int16 `json:"max" validate:"required,min=0,max=2000"`
}

type PlaceTranslationRequest struct {
	Title       string                     `json:"title" validate:"required,min=1,max=255"`
	Description string                     `json:"description" validate:"required,min=1,max=255"`
	Locale      string                     `json:"locale" validate:"required,min=1,max=255"`
	MarkdownURL string                     `json:"markdownUrl" validate:"required,min=1,max=255"`
	Seo         PlaceTranslationSeoRequest `json:"seo" validate:"required,dive,required"`
}

type PlaceTranslationSeoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1,max=255"`
	Keywords    string `json:"keywords" validate:"required,min=1,max=255"`
}

type PlaceRestorationRequest struct {
	StartDate string `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate" validate:"omitempty,datetime=2006-01-02"`
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
