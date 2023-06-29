package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type FeatureDeleteCommand struct{}

type FeatureDeleteResult struct{}

type FeatureDeleteHandler decorator.CommandHandler[FeatureDeleteCommand, *FeatureDeleteResult]

type featureDeleteHandler struct {
	repo    feature.Repository
	factory feature.Factory
}

type FeatureDeleteHandlerConfig struct {
	Repo     feature.Repository
	Factory  feature.Factory
	CqrsBase decorator.Base
}

func NewFeatureDeleteHandler(config FeatureDeleteHandlerConfig) FeatureDeleteHandler {
	return decorator.ApplyCommandDecorators[FeatureDeleteCommand, *FeatureDeleteResult](
		featureDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h featureDeleteHandler) Handle(ctx context.Context, command FeatureDeleteCommand) (*FeatureDeleteResult, *i18np.Error) {
	return &FeatureDeleteResult{}, nil
}
