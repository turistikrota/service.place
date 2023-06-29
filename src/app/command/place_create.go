package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlaceCreateCommand struct{}

type PlaceCreateResult struct{}

type PlaceCreateHandler decorator.CommandHandler[PlaceCreateCommand, *PlaceCreateResult]

type placeCreateHandler struct {
	repo    place.Repository
	factory place.Factory
}

type PlaceCreateHandlerConfig struct {
	Repo     place.Repository
	Factory  place.Factory
	CqrsBase decorator.Base
}

func NewPlaceCreateHandler(config PlaceCreateHandlerConfig) PlaceCreateHandler {
	return decorator.ApplyCommandDecorators[PlaceCreateCommand, *PlaceCreateResult](
		placeCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeCreateHandler) Handle(ctx context.Context, command PlaceCreateCommand) (*PlaceCreateResult, *i18np.Error) {
	return &PlaceCreateResult{}, nil
}
