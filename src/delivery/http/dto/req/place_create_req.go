package req

import (
	"strings"

	"github.com/ssibrahimbas/slug"
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/domain/place"
)

type PlaceCreateRequest struct {
	FeatureUUIDs []string                  `json:"featureUUIDs" validate:"required,min=1,max=10,dive,object_id"`
	Images       []PlaceImageRequest       `json:"images" validate:"required,min=1,max=10,dive,required"`
	Coordinates  []float64                 `json:"coordinates" validate:"required,min=2,max=2,dive,required"`
	TimeSpent    PlaceTimeSpentRequest     `json:"timeSpent" validate:"required,dive,required"`
	Translations []PlaceTranslationRequest `json:"translations" validate:"required,min=1,max=3,dive,required"`
	IsPayed      *bool                     `json:"isPayed" validate:"required"`
	Type         string                    `json:"type" validate:"required,oneof=eating coffee bar beach amaze shopping transport culture nature health sport nightlife other"`
}

func (r *PlaceCreateRequest) ToCommand() command.PlaceCreateCommand {
	return command.PlaceCreateCommand{
		FeatureUUIDs:     r.FeatureUUIDs,
		Images:           r.toImages(),
		Translations:     r.toTranslations(),
		AverageTimeSpent: r.toAverageTimeSpent(),
		Coordinates:      r.Coordinates,
		IsPayed:          r.toIsPayed(),
		Type:             place.Type(r.Type),
	}
}

func (r *PlaceCreateRequest) toImages() []place.Image {
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

func (r *PlaceCreateRequest) toTranslations() map[place.Locale]place.Translations {
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

func (r *PlaceCreateRequest) toAverageTimeSpent() place.TimeSpent {
	return place.TimeSpent{
		Min: r.TimeSpent.Min,
		Max: r.TimeSpent.Max,
	}
}

func (r *PlaceCreateRequest) toIsPayed() bool {
	if r.IsPayed == nil {
		return false
	}
	return *r.IsPayed
}
