package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	FeatureUpdateCommand struct {
		UUID         string
		Icon         string
		Translations map[feature.Locale]feature.Translations
	}
	FeatureUpdateResult  struct{}
	FeatureUpdateHandler decorator.CommandHandler[FeatureUpdateCommand, *FeatureUpdateResult]
	featureUpdateHandler struct {
		repo    feature.Repository
		factory feature.Factory
	}
	FeatureUpdateHandlerConfig struct {
		Repo     feature.Repository
		Factory  feature.Factory
		CqrsBase decorator.Base
	}
)

func NewFeatureUpdateHandler(config FeatureUpdateHandlerConfig) FeatureUpdateHandler {
	return decorator.ApplyCommandDecorators[FeatureUpdateCommand, *FeatureUpdateResult](
		featureUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h featureUpdateHandler) Handle(ctx context.Context, command FeatureUpdateCommand) (*FeatureUpdateResult, *i18np.Error) {
	f := h.factory.New(command.Icon, command.Translations)
	err := h.repo.Update(ctx, command.UUID, f)
	if err != nil {
		return nil, err
	}
	return &FeatureUpdateResult{}, nil
}
