package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
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
