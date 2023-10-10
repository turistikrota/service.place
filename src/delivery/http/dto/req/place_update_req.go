package req

import (
	"strings"
	"time"

	"github.com/ssibrahimbas/slug"
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/domain/place"
)

type PlaceUpdateRequest struct {
	UUID         string
	FeatureUUIDs []string                  `json:"featureUUIDs" validate:"required,min=1,max=10,dive,object_id"`
	Images       []PlaceImageRequest       `json:"images" validate:"required,min=1,max=10,dive,required"`
	Coordinates  []float64                 `json:"coordinates" validate:"required,min=2,max=2,dive,required"`
	TimeSpent    PlaceTimeSpentRequest     `json:"timeSpent" validate:"required,dive,required"`
	Translations []PlaceTranslationRequest `json:"translations" validate:"required,min=1,max=3,dive,required"`
	Restorations []PlaceRestorationRequest `json:"restorations" validate:"omitempty,min=0,max=100,dive,required"`
	IsPayed      *bool                     `json:"isPayed" validate:"required"`
	Type         string                    `json:"type" validate:"required,oneof=eating coffee bar beach amaze shopping transport culture nature health sport nightlife other"`
}

func (r *PlaceUpdateRequest) Load(req *PlaceDetailRequest) *PlaceUpdateRequest {
	r.UUID = req.UUID
	return r
}

func (r *PlaceUpdateRequest) ToCommand() command.PlaceUpdateCommand {
	return command.PlaceUpdateCommand{
		FeatureUUIDs:     r.FeatureUUIDs,
		Images:           r.toImages(),
		Translations:     r.toTranslations(),
		AverageTimeSpent: r.toAverageTimeSpent(),
		Restorations:     r.toRestorations(),
		Coordinates:      r.Coordinates,
		IsPayed:          r.toIsPayed(),
		UUID:             r.UUID,
		Type:             place.Type(r.Type),
	}
}

func (r *PlaceUpdateRequest) toRestorations() []place.Restoration {
	restorations := make([]place.Restoration, len(r.Restorations))
	for i, restoration := range r.Restorations {
		startDate, _ := time.Parse("2006-01-02", restoration.StartDate)
		endDate, _ := time.Parse("2006-01-02", restoration.EndDate)
		restorations[i] = place.Restoration{
			StartDate: &startDate,
			EndDate:   &endDate,
		}
	}
	return restorations
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
