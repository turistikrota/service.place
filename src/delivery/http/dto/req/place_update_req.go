package req

import (
	"strings"

	"api.turistikrota.com/place/src/app/command"
	"api.turistikrota.com/place/src/domain/place"
	"github.com/ssibrahimbas/slug"
)

type PlaceUpdateRequest struct {
	FeatureUUIDs []string                  `json:"feature_uuids" validate:"required,min=1,max=10,dive,object_id"`
	Images       []PlaceImageRequest       `json:"images" validate:"required,min=1,max=10,dive,required"`
	Coordinates  []float64                 `json:"coordinates" validate:"required,min=2,max=2,dive,required"`
	TimeSpent    PlaceTimeSpentRequest     `json:"time_spent" validate:"required,dive,required"`
	Translations []PlaceTranslationRequest `json:"translations" validate:"required,min=1,max=3,dive,required"`
	IsPayed      *bool                     `json:"is_payed" validate:"required"`
}

func (r *PlaceUpdateRequest) ToCommand() command.PlaceUpdateCommand {
	return command.PlaceUpdateCommand{
		FeatureUUIDs:     r.FeatureUUIDs,
		Images:           r.toImages(),
		Translations:     r.toTranslations(),
		AverageTimeSpent: r.toAverageTimeSpent(),
		Coordinates:      r.Coordinates,
		IsPayed:          r.toIsPayed(),
	}
}

func (r *PlaceUpdateRequest) toImages() []place.Image {
	images := make([]place.Image, len(r.Images))
	for i, image := range r.Images {
		var order int16
		if image.Order != nil {
			order = *image.Order
		}
		images[i] = place.Image{
			Url:   image.Url,
			Order: order,
		}
	}
	return images
}

func (r *PlaceUpdateRequest) toTranslations() map[place.Locale]place.Translations {
	translations := make(map[place.Locale]place.Translations)
	for _, translation := range r.Translations {
		translations[place.Locale(translation.Locale)] = place.Translations{
			Title:       translation.Title,
			Description: translation.Description,
			Slug:        slug.New(translation.Title, slug.Lang(strings.ToUpper(translation.Locale))),
			MarkdownURL: translation.MarkdownURL,
			Seo:         translation.Seo.ToSeo(),
		}
	}
	return translations
}

func (r *PlaceUpdateRequest) toAverageTimeSpent() place.TimeSpent {
	return place.TimeSpent{
		Min: r.TimeSpent.Min,
		Max: r.TimeSpent.Max,
	}
}

func (r *PlaceUpdateRequest) toIsPayed() bool {
	if r.IsPayed == nil {
		return false
	}
	return *r.IsPayed
}
