package req

import (
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type FeatureUpdateRequest struct {
	UUID         string
	Icon         string                      `json:"icon" validate:"required,min=1,max=255"`
	Translations []FeatureTranslationRequest `json:"translations" validate:"min=1,max=3,dive,required"`
}

func (r *FeatureUpdateRequest) Load(req *FeatureDetailRequest) *FeatureUpdateRequest {
	r.UUID = req.UUID
	return r
}

func (r *FeatureUpdateRequest) ToCommand() command.FeatureUpdateCommand {
	return command.FeatureUpdateCommand{
		UUID:         r.UUID,
		Icon:         r.Icon,
		Translations: r.toTranslations(),
	}
}

func (r *FeatureUpdateRequest) toTranslations() map[feature.Locale]feature.Translations {
	translations := make(map[feature.Locale]feature.Translations)
	for _, translation := range r.Translations {
		translations[feature.Locale(translation.Locale)] = feature.Translations{
			Title:       translation.Title,
			Description: translation.Description,
		}
	}
	return translations
}
