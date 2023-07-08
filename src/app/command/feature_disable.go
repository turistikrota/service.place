package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type (
	FeatureDisableCommand struct {
		UUID string
	}
	FeatureDisableResult  struct{}
	FeatureDisableHandler decorator.CommandHandler[FeatureDisableCommand, *FeatureDisableResult]
	featureDisableHandler struct {
		repo    feature.Repository
		factory feature.Factory
	}
	FeatureDisableHandlerConfig struct {
		Repo     feature.Repository
		Factory  feature.Factory
		CqrsBase decorator.Base
	}
)

func NewFeatureDisableHandler(config FeatureDisableHandlerConfig) FeatureDisableHandler {
	return decorator.ApplyCommandDecorators[FeatureDisableCommand, *FeatureDisableResult](
		featureDisableHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h featureDisableHandler) Handle(ctx context.Context, command FeatureDisableCommand) (*FeatureDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &FeatureDisableResult{}, nil
}
