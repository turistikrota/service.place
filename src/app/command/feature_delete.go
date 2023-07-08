package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
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
