package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceUpdateCommand struct{}
	PlaceUpdateResult  struct{}
	PlaceUpdateHandler decorator.CommandHandler[PlaceUpdateCommand, *PlaceUpdateResult]
	placeUpdateHandler struct {
		repo    place.Repository
		factory place.Factory
	}
	PlaceUpdateHandlerConfig struct {
		Repo     place.Repository
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceUpdateHandler(config PlaceUpdateHandlerConfig) PlaceUpdateHandler {
	return decorator.ApplyCommandDecorators[PlaceUpdateCommand, *PlaceUpdateResult](
		placeUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeUpdateHandler) Handle(ctx context.Context, command PlaceUpdateCommand) (*PlaceUpdateResult, *i18np.Error) {
	return &PlaceUpdateResult{}, nil
}
