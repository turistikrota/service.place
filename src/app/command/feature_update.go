package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type FeatureUpdateCommand struct{}

type FeatureUpdateResult struct{}

type FeatureUpdateHandler decorator.CommandHandler[FeatureUpdateCommand, *FeatureUpdateResult]

type featureUpdateHandler struct {
	repo    feature.Repository
	factory feature.Factory
}

type FeatureUpdateHandlerConfig struct {
	Repo     feature.Repository
	Factory  feature.Factory
	CqrsBase decorator.Base
}

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
	return &FeatureUpdateResult{}, nil
}
