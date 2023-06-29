package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceDisableCommand struct {
		UUID string
	}
	PlaceDisableResult  struct{}
	PlaceDisableHandler decorator.CommandHandler[PlaceDisableCommand, *PlaceDisableResult]
	placeDisableHandler struct {
		repo    place.Repository
		factory place.Factory
	}
	PlaceDisableHandlerConfig struct {
		Repo     place.Repository
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceDisableHandler(config PlaceDisableHandlerConfig) PlaceDisableHandler {
	return decorator.ApplyCommandDecorators[PlaceDisableCommand, *PlaceDisableResult](
		placeDisableHandler{
			repo:    config.Repo,
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
