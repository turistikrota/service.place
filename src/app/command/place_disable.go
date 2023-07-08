package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	PlaceDisableCommand struct {
		UUID string
	}
	PlaceDisableResult  struct{}
	PlaceDisableHandler decorator.CommandHandler[PlaceDisableCommand, *PlaceDisableResult]
	placeDisableHandler struct {
		repo    place.Repository
		events  place.Events
		factory place.Factory
	}
	PlaceDisableHandlerConfig struct {
		Repo     place.Repository
		Events   place.Events
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceDisableHandler(config PlaceDisableHandlerConfig) PlaceDisableHandler {
	return decorator.ApplyCommandDecorators[PlaceDisableCommand, *PlaceDisableResult](
		placeDisableHandler{
			repo:    config.Repo,
			events:  config.Events,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeDisableHandler) Handle(ctx context.Context, command PlaceDisableCommand) (*PlaceDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &PlaceDisableResult{}, nil
}
