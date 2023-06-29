package req

import (
	"api.turistikrota.com/place/src/app/command"
	"api.turistikrota.com/place/src/domain/feature"
)

type FeatureUpdateRequest struct {
	Icon         string                      `json:"icon" validate:"required,min=1,max=255"`
	Translations []FeatureTranslationRequest `json:"translations" validate:"min=1,max=3,dive,required"`
}

func (r *FeatureUpdateRequest) ToCommand() command.FeatureUpdateCommand {
	return command.FeatureUpdateCommand{
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