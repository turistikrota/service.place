package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	PlaceEnableCommand struct {
		UUID string
	}
	PlaceEnableResult  struct{}
	PlaceEnableHandler decorator.CommandHandler[PlaceEnableCommand, *PlaceEnableResult]
	placeEnableHandler struct {
		repo    place.Repository
		events  place.Events
		factory place.Factory
	}
	PlaceEnableHandlerConfig struct {
		Repo     place.Repository
		Events   place.Events
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceEnableHandler(config PlaceEnableHandlerConfig) PlaceEnableHandler {
	return decorator.ApplyCommandDecorators[PlaceEnableCommand, *PlaceEnableResult](
		placeEnableHandler{
			repo:    config.Repo,
			events:  config.Events,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeEnableHandler) Handle(ctx context.Context, command PlaceEnableCommand) (*PlaceEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &PlaceEnableResult{}, nil
}
