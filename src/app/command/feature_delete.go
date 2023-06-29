package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	FeatureDeleteCommand struct {
		UUID string
	}
	FeatureDeleteResult  struct{}
	FeatureDeleteHandler decorator.CommandHandler[FeatureDeleteCommand, *FeatureDeleteResult]
	featureDeleteHandler struct {
		repo    feature.Repository
		factory feature.Factory
	}
	FeatureDeleteHandlerConfig struct {
		Repo     feature.Repository
		Factory  feature.Factory
		CqrsBase decorator.Base
	}
)

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
	err := h.repo.Delete(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &FeatureDeleteResult{}, nil
}
