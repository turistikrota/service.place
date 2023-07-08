package req

import (
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type FeatureCreateRequest struct {
	Icon         string                      `json:"icon" validate:"required,min=1,max=255"`
	Translations []FeatureTranslationRequest `json:"translations" validate:"min=1,max=3,dive,required"`
}

type FeatureTranslationRequest struct {
	Locale      string `json:"locale" validate:"required,min=1,max=255"`
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1,max=255"`
}

func (r *FeatureCreateRequest) ToCommand() command.FeatureCreateCommand {
	return command.FeatureCreateCommand{
		Icon:         r.Icon,
		Translations: r.toTranslations(),
	}
}

func (r *FeatureCreateRequest) toTranslations() map[feature.Locale]feature.Translations {
	translations := make(map[feature.Locale]feature.Translations)
	for _, translation := range r.Translations {
		translations[feature.Locale(translation.Locale)] = feature.Translations{
			Title:       translation.Title,
			Description: translation.Description,
		}
	}
	return translations
}
