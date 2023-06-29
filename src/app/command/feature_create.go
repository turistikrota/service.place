package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type FeatureCreateCommand struct{}

type FeatureCreateResult struct{}

type FeatureCreateHandler decorator.CommandHandler[FeatureCreateCommand, *FeatureCreateResult]

type featureCreateHandler struct {
	repo    feature.Repository
	factory feature.Factory
}

type FeatureCreateHandlerConfig struct {
	Repo     feature.Repository
	Factory  feature.Factory
	CqrsBase decorator.Base
}

func NewFeatureCreateHandler(config FeatureCreateHandlerConfig) FeatureCreateHandler {
	return decorator.ApplyCommandDecorators[FeatureCreateCommand, *FeatureCreateResult](
		featureCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h featureCreateHandler) Handle(ctx context.Context, command FeatureCreateCommand) (*FeatureCreateResult, *i18np.Error) {
	return &FeatureCreateResult{}, nil
}
