package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type (
	FeatureCreateCommand struct {
		Icon         string
		Translations map[feature.Locale]feature.Translations
	}
	FeatureCreateResult  struct{}
	FeatureCreateHandler decorator.CommandHandler[FeatureCreateCommand, *FeatureCreateResult]
	featureCreateHandler struct {
		repo    feature.Repository
		factory feature.Factory
	}
	FeatureCreateHandlerConfig struct {
		Repo     feature.Repository
		Factory  feature.Factory
		CqrsBase decorator.Base
	}
)

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
	f := h.factory.New(command.Icon, command.Translations)
	err := h.repo.Create(ctx, f)
	if err != nil {
		return nil, err
	}
	return &FeatureCreateResult{}, nil
}
