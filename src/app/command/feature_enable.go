package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	FeatureEnableCommand struct {
		UUID string
	}
	FeatureEnableResult  struct{}
	FeatureEnableHandler decorator.CommandHandler[FeatureEnableCommand, *FeatureEnableResult]
	featureEnableHandler struct {
		repo    feature.Repository
		factory feature.Factory
	}
	FeatureEnableHandlerConfig struct {
		Repo     feature.Repository
		Factory  feature.Factory
		CqrsBase decorator.Base
	}
)

func NewFeatureEnableHandler(config FeatureEnableHandlerConfig) FeatureEnableHandler {
	return decorator.ApplyCommandDecorators[FeatureEnableCommand, *FeatureEnableResult](
		featureEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h featureEnableHandler) Handle(ctx context.Context, command FeatureEnableCommand) (*FeatureEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &FeatureEnableResult{}, nil
}
